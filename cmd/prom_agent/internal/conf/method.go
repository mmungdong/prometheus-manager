package conf

func (c *Bootstrap) GetFlag() string {
	if c == nil {
		return ""
	}
	return c.flag
}

func (c *Bootstrap) GetServer() *Server {
	if c == nil {
		return nil
	}
	return c.Server
}

func (c *Bootstrap) GetKafka() *Kafka {
	if c == nil {
		return nil
	}

	return c.Kafka
}

func (c *Bootstrap) GetAlert() *Alert {
	if c == nil {
		return nil
	}
	return c.Alert
}

func (c *Bootstrap) GetSyncStrategy() *SyncStrategy {
	if c == nil {
		return nil
	}
	return c.SyncStrategy
}

func (c *Server) GetHttp() *Http {
	if c == nil {
		return nil
	}
	return c.Http
}

func (c *Server) GetName() string {
	if c == nil {
		return ""
	}
	return c.Name
}

func (c *Server) GetMode() string {
	if c == nil {
		return ""
	}
	return c.Mode
}

func (c *Server) GetMetadata() map[string]string {
	if c == nil {
		return nil
	}
	return c.Metadata
}

func (c *Http) GetHost() string {
	if c == nil {
		return ""
	}
	return c.Host
}

func (c *Http) GetPort() int {
	if c == nil {
		return 0
	}
	return c.Port
}

func (c *Kafka) GetEnable() bool {
	if c == nil {
		return false
	}
	return c.Enable
}

func (c *Kafka) GetEndpoints() []string {
	if c == nil {
		return nil
	}
	return c.Endpoints
}

func (c *Alert) GetTopic() string {
	if c == nil {
		return ""
	}
	return c.Topic
}

func (c *SyncStrategy) GetEnable() bool {
	if c == nil {
		return false
	}
	return c.Enable
}

func (c *SyncStrategy) GetTopic() string {
	if c == nil {
		return ""
	}
	return c.Topic
}
