package agent

import (
	"encoding/json"
	"fmt"
	"time"

	client "github.com/Gimulator/client-go"
	uuid "github.com/satori/go.uuid"
)

var name = "random-agent"

const (
	typeVerdict  = "verdict"
	typeAction   = "action"
	typeRegister = "register"
	namespace    = "paper-soccer"
	worldName    = "world"

	apiTimeWait = 3
)

type controller struct {
	*client.Client
}

func newController(ch chan client.Object) (controller, error) {
	name = name + "-" + uuid.NewV4().String()
	fmt.Println(name)

	cli, err := client.NewClient(ch)
	if err != nil {
		return controller{}, err
	}

	err = cli.Watch(client.Key{
		Name:      worldName,
		Namespace: namespace,
		Type:      typeVerdict,
	})

	if err != nil {
		return controller{}, err
	}

	return controller{
		cli,
	}, nil
}

func (c *controller) setAction(move Move) error {
	fmt.Println("action", move)
	val, err := json.Marshal(move)
	if err != nil {
		return err
	}

	obj := client.Object{
		Key: client.Key{
			Type:      typeAction,
			Namespace: namespace,
			Name:      name,
		},
		Value: string(val),
	}

	for {
		err = c.Set(obj)
		if err == nil {
			return nil
		}

		time.Sleep(time.Second * apiTimeWait)
	}
}

func (c *controller) setRegister() {
	obj := client.Object{
		Key: client.Key{
			Type:      typeRegister,
			Namespace: namespace,
			Name:      name,
		},
		Value: "",
	}

	for {
		err := c.Set(obj)
		if err == nil {
			break
		}

		time.Sleep(time.Second * apiTimeWait)
	}
}
