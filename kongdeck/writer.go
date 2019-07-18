package kongdeck

import (
	"github.com/davidcv5/kongup/kongfig"
	deck "github.com/hbagdi/deck/file"
)

// KongfigToDeck maps kongfig to deck and saves it to a file
func KongfigToDeck(config *kongfig.Config, tags []string, filename string) error {
	file, err := fromKongfig(config)
	state, _, err := deck.GetStateFromContent(file)
	if err != nil {
		return err
	}
	if err = deck.KongStateToFile(state, tags, filename); err != nil {
		return err
	}
	return nil
	// c, err := yaml.Marshal(file)
	// if err != nil {
	// 	return err
	// }

	// if filename == "-" {
	// 	_, err = fmt.Print(string(c))
	// } else {
	// 	err = ioutil.WriteFile(filename, c, 0600)
	// }

	// if err != nil {
	// 	return err
	// }
	// return nil
}
