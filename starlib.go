package starlib

import (
	"fmt"

	"github.com/drudge/starlib/bsoup"
	"github.com/drudge/starlib/compress/gzip"
	"github.com/drudge/starlib/dataframe"
	"github.com/drudge/starlib/encoding/base64"
	"github.com/drudge/starlib/encoding/csv"
	"github.com/drudge/starlib/encoding/json"
	"github.com/drudge/starlib/encoding/yaml"
	"github.com/drudge/starlib/geo"
	"github.com/drudge/starlib/hash"
	"github.com/drudge/starlib/hmac"
	"github.com/drudge/starlib/html"
	"github.com/drudge/starlib/http"
	"github.com/drudge/starlib/math"
	"github.com/drudge/starlib/re"
	"github.com/drudge/starlib/time"
	"github.com/drudge/starlib/xlsx"
	"github.com/drudge/starlib/zipfile"
	"go.starlark.net/starlark"
)

// Version is the current semver for the entire starlib library
const Version = "0.5.0"

// Loader presents the starlib library as a loader
func Loader(thread *starlark.Thread, module string) (dict starlark.StringDict, err error) {
	switch module {
	case time.ModuleName:
		return starlark.StringDict{"time": time.Module}, nil
	case gzip.ModuleName:
		return starlark.StringDict{"gzip": gzip.Module}, nil
	case http.ModuleName:
		return http.LoadModule()
	case xlsx.ModuleName:
		return xlsx.LoadModule()
	case html.ModuleName:
		return html.LoadModule()
	case bsoup.ModuleName:
		return bsoup.LoadModule()
	case zipfile.ModuleName:
		return zipfile.LoadModule()
	case re.ModuleName:
		return re.LoadModule()
	case base64.ModuleName:
		return base64.LoadModule()
	case csv.ModuleName:
		return csv.LoadModule()
	case json.ModuleName:
		return starlark.StringDict{"json": json.Module}, nil
	case yaml.ModuleName:
		return yaml.LoadModule()
	case geo.ModuleName:
		return geo.LoadModule()
	case math.ModuleName:
		return starlark.StringDict{"math": math.Module}, nil
	case hash.ModuleName:
		return hash.LoadModule()
	case hmac.ModuleName:
		return hmac.LoadModule()
	case dataframe.ModuleName:
		return starlark.StringDict{"dataframe": dataframe.Module}, nil
	}

	return nil, fmt.Errorf("invalid module %q", module)
}
