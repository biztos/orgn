/* binsanity.go - auto-generated; edit at your own peril!

More info: https://github.com/biztos/binsanity

*/

package orgn

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"errors"
	"io"
	"sort"
)

// Asset returns the byte content of the asset for the given name, or an error
// if no such asset is available.
func Asset(name string) ([]byte, error) {

	_, found := binsanity_cache[name]
	if !found {
		i := sort.SearchStrings(binsanity_names, name)
		if i == len(binsanity_names) || binsanity_names[i] != name {
			return nil, errors.New("Asset not found.")
		}

		// We ignore errors because we controlled the data from the begining.
		// It's not perfect but it seems better than having additional funcs
		// hanging around that might confuse the user: tried that already, not
		// nicer.
		decoded, _ := base64.StdEncoding.DecodeString(binsanity_data[i])
		buf := bytes.NewReader(decoded)
		gzr, _ := gzip.NewReader(buf)
		defer gzr.Close()
		data, _ := io.ReadAll(gzr)

		// Not cached, so decode and cache it.
		binsanity_cache[name] = data

	}
	return binsanity_cache[name], nil

}

// MustAsset returns the byte content of the asset for the given name, or
// panics if no such asset is available.
func MustAsset(name string) []byte {
	b, err := Asset(name)
	if err != nil {
		panic(err.Error())
	}
	return b
}

// MustAssetString returns the string content of the asset for the given name,
// or panics if no such asset is available.  This is a convenience function
// for string(MustAsset(name)).
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetNames returns the sorted names of the assets.
func AssetNames() []string {
	return binsanity_names
}

// this must remain sorted or everything breaks!
var binsanity_names = []string{
	"default_page.html",
}

// only decode once per asset.
var binsanity_cache = map[string][]byte{}

// assets are gzipped and base64 encoded
var binsanity_data = []string{
	"H4sIAAAAAAAA/7JRTMlPLqksSFXIKMnNseOyyU0tSVRIzkgsKk4tsVUqLUmzULLjsskwtPNIzcnJVwjPL8pJsdHPMLQDBAAA///LRWsSOgAAAA==",
}
