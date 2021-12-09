package main

import "encoding/json"

type MetadataServer struct {
	CodigoServidor string `json:"codigo_servidor"`
	Signature      []byte `json:"signature"`
}

func (m *MetadataServer) ToJson() ([]byte, error) {
	return json.Marshal(m)
}
