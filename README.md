# `package s2covergeojson`

```
github.com/ponlawat-w/golang-s2covergeojson
```

```go
// S2CellIDToBase64 converts s2 CellID into base64 representation
func S2CellIDToBase64(cellID s2.CellID) string 

// Cover function returns s2 cell union from specified geojson feature and RegionCoverer
func Cover(feature *geojson.Feature, regionCoverer s2.RegionCoverer) (s2.CellUnion, error)

// ReadGeoJSON reads geojson file from specified path, and return parsed geojson.FeatureCollection
func ReadGeoJSON(filePath string) (*geojson.FeatureCollection, error)

// WriteTokensToFile writes list of tokens in cell union, one per line
func WriteTokensToFile(cellUnion s2.CellUnion, file *os.File) error

// WriteBase64TokensToFile writes list of base64 tokens in cell union, one per line
func WriteBase64TokensToFile(cellUnion s2.CellUnion, file *os.File) error

// WriteTokensToPath writes list of tokens in cell union, one per line
func WriteTokensToPath(cellUnion s2.CellUnion, filePath string) error

// WriteBase64TokensToPath writes list of base64 tokens in cell union, one per line
func WriteBase64TokensToPath(cellUnion s2.CellUnion, filePath string) error
```
