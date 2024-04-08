package rpc

import (
	"context"

	"github.com/off-chain-storage/GoSphere/go-sphere/rpc/iface"
	"github.com/off-chain-storage/GoSphere/message"
)

type router struct {
	routerClient iface.RouterClient
	cs           *ClientService
}

func (r *router) SendDataToPropagationManager(ctx context.Context, blockData []byte) error {

	blk, err := message.BuildData(blockData)
	if err != nil {
		log.Error("Failed to build block data")
		return err
	}

	// Send Acknowledgement for time measurement - temp service
	go r.cs.SendUDPMessage("Data sent to P-M from C-R")

	_, err = r.routerClient.SendDataToPropagationManager(ctx, blk)
	if err != nil {
		log.Error("Failed to send data to propagation manager")
		return err
	}

	log.Info("Data sent to propagation manager successfully")

	return nil
}
