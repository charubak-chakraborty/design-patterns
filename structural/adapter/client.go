package adapter

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	com.insertIntoLightningPort()
}
