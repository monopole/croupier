// System and game logic.

// An instance of GameManager is a peer to other instances on the net.
// Each has an embedded V23 service, and is a direct client to the V23
// services held by all the other instances.  It finds all the other
// instances, figures out what it should call itself, and fires off
// go routines to manage data coming in on various channels.
//
// The GameManager is presumably owned by whatever owns the UX event
// loop,
//
// During play, UX or underlying android/iOS events may trigger calls
// to other V23 services, Likewise, an incoming RPC may change data
// held by the manager, to ultimately impact the UX (e.g. a card is
// passed in by another player).

package game

import (
	"github.com/monopole/croupier/ifc"
	"github.com/monopole/croupier/service"
	"log"
	"strconv"
	"time"
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/naming"
	"v.io/v23/options"
	_ "v.io/x/ref/runtime/factories/generic"
)

const rootName = "croupier/player"
const namespaceRoot = "/104.197.96.113:3389"

// The number of instances of this program to run in a demo.
// Need an exact count to wire them up properly.
const expectedInstances = 2

func serverName(n int) string {
	return rootName + fmt.Sprintf("%04d", n)
}

type GameManager struct {
	ctx      *context.T
	myNumber int // my player number
	master   ifc.GameBuddyClientStub
	chatty   bool    // If true, send fortunes back and forth and log them.  For fun.
	originX  float32 // remember where to put the card
	originY  float32
}

func (gm *GameManager) MyNumber() int {
	return gm.myNumber
}

func (gm *GameManager) GetOriginX() float32 {
	return gm.originX
}

func (gm *GameManager) GetOriginY() float32 {
	return gm.originY
}

func (gm *GameManager) SetOrigin(x, y float32) {
	gm.originX = x
	gm.originY = y
}

func NewGameManager(ctx *context.T) *GameManager {
	gm := &GameManager{ctx, 0, nil, true}
	gm.initialize()
	return gm
}

func (gm *GameManager) initialize() {
	v23.GetNamespace(gm.ctx).SetRoots(namespaceRoot)

	gm.myNumber = gm.playerCount()

	// If there are no players, I register as player 1.  If there is one
	// player already, I register as player 2, etc.
	gm.registerService()

	// No matter who I am, I am a client to server0.
	gm.master = ifc.GameBuddyClient(serverName(0))
}

// Scan mounttable for count of services matching "{rootName}*"
func (gm *GameManager) playerCount() (count int) {
	count = 0
	rCtx, cancel := context.WithTimeout(gm.ctx, time.Minute)
	defer cancel()
	ns := v23.GetNamespace(rCtx)
	pattern := rootName + "*"
	c, err := ns.Glob(rCtx, pattern)
	if err != nil {
		log.Printf("ns.Glob(%q) failed: %v", pattern, err)
		return
	}
	for res := range c {
		switch v := res.(type) {
		case *naming.GlobReplyEntry:
			if v.Value.Name != "" {
				count++
				if gm.chatty {
					log.Println(v.Value.Name)
				}
			}
		default:
		}
	}
	return
}

// Register a service in the namespace and begin serving.
func (gm *GameManager) registerService() {
	s := MakeServer(gm.ctx)
	myName := serverName(gm.myNumber)
	log.Printf("Calling myself %s\n", myName)
	err := s.Serve(myName, ifc.GameBuddyServer(service.Make()), MakeAuthorizer())
	if err != nil {
		log.Panic("Error serving service: ", err)
	}
}

func (gm *GameManager) WhoHasTheCard() int {
	who, _ := gm.master.WhoHasCard(gm.ctx, options.SkipServerEndpointAuthorization{})
	return int(who)
}

// Modify the game state, and send it to all players, starting with the
// player that's gonna get the card.
func (gm *GameManager) PassTheCard() {
	if gm.chatty {
		log.Printf("Sending to %v\n", serverName(gm.myNumber)+" "+time.Now().String())
	}
	if err := gm.master.SendCardTo(gm.ctx, int32((gm.myNumber+1)%expectedInstances),
		options.SkipServerEndpointAuthorization{}); err != nil {
		log.Printf("error sending card: %v\n", err)
	}
}
