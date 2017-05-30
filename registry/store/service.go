// Copyright 2016 IBM Corporation
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package store

import "fmt"

const (
	DefaultAddress = "0.0.0.0"
	DefaultPortNumber = 80
	DefaultProtocol   = "TCP"
)

// Service represents a runtime service group.
type Service struct {
	ServiceName  string
	Address      string
	Ports        PortList
	ExternalName string
}

type Port struct {
	Name     string
	Port     int
	Protocol string
}

type PortList []*Port

func (s *Service) DeepClone() *Service {
	cloned := *s
	if s.Address == "" {
		cloned.Address = DefaultAddress
	}

	if s.Ports == nil {
		s.Ports = make(PortList, 0)
	}
	cloned.Ports = make(PortList, len(s.Ports))
	for idx, port := range s.Ports {
		cloned.Ports[idx] = port.DeepClone()
	}

	return &cloned
}

func (p *Port) DeepClone() *Port {
	cloned := *p
	if p.Port == 0 {
		cloned.Port = DefaultPortNumber
	}
	if p.Protocol == "" {
		cloned.Protocol = DefaultProtocol
	}
	return &cloned
}

func (s *Service) PutPort(port *Port) {
	s.Ports = append(s.Ports, port)
}

func (pl PortList) Get(portName string) (*Port, bool) {
	for _, port := range pl {
		if port.Name == portName {
			return port, true
		}
	}

	return nil, false
}

func (p *Port) Equals(port *Port) bool {
	return p.Port == port.Port && p.Protocol == port.Protocol
}

func (s *Service) String() string {
	return fmt.Sprintf("%s", s.ServiceName)
}
