package client

import (
	"context"
	"encoding/json"
	"errors"
)

func (this *Client) deleteCharacter() error {
	var data struct {
		Name string
	}

	err := json.Unmarshal([]byte(this.body), &data)

	if err != nil {
		return errors.New("Client.deleteCharacter: failed to unmarshal")
	}

	db, _ := this.manager.DB.Acquire(context.TODO())
	defer db.Release()

	const query = "DELETE FROM characters WHERE name=$1 AND user_id=$2;"
	t, err := db.Exec(context.TODO(), query, data.Name, this.user.ID)

	if err != nil || t.RowsAffected() == 0 {
		this.sendMessage(&Dictionary{"error": "failed to delete character"})
		return err
	}
	this.sendMessage(&Dictionary{"success": "character deleted"})
	return nil
}