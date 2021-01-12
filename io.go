package s2covergeojson

import (
	"io/ioutil"
	"os"

	"github.com/golang/geo/s2"
	"github.com/paulmach/go.geojson"
)

// ReadGeoJSON reads geojson file from specified path, and return parsed geojson.FeatureCollection
func ReadGeoJSON(filePath string) (*geojson.FeatureCollection, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return geojson.UnmarshalFeatureCollection(data);
}

// WriteTokensToFile writes list of tokens in cell union, one per line
func WriteTokensToFile(cellUnion s2.CellUnion, file *os.File) error {
	for _, cell := range cellUnion {
		if _, err := file.Write([]byte(cell.ToToken() + "\n")); err != nil {
			return err
		}
	}

	return nil;
}

// WriteBase64TokensToFile writes list of base64 tokens in cell union, one per line
func WriteBase64TokensToFile(cellUnion s2.CellUnion, file *os.File) error {
	for _, cell := range cellUnion {
		if _, err := file.Write([]byte(S2CellIDToBase64(cell) + "\n")); err != nil {
			return err
		}
	}

	return nil;
}

// WriteTokensToPath writes list of tokens in cell union, one per line
func WriteTokensToPath(cellUnion s2.CellUnion, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	return WriteTokensToFile(cellUnion, file)
}

// WriteBase64TokensToPath writes list of base64 tokens in cell union, one per line
func WriteBase64TokensToPath(cellUnion s2.CellUnion, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	return WriteBase64TokensToFile(cellUnion, file)
}
