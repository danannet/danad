package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/danannet/danad/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.DanadMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.DanadMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.DanadMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.DanadMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.DanadMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.DanadMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.DanadMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.DanadMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.DanadMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.DanadMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.DanadMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.DanadMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.DanadMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.DanadMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.DanadMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.DanadMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.DanadMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.DanadMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.DanadMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.DanadMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.DanadMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.DanadMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.DanadMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.DanadMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.DanadMessage_BanRequest{}),
	reflect.TypeOf(protowire.DanadMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
