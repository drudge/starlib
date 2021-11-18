package dataframe

import (
	"fmt"

	"go.starlark.net/starlark"
)

type dataframeAttrImpl func(*DataFrame) (starlark.Value, error)

var dataframeAttributes = map[string]dataframeAttrImpl{
	"at":      dataframeAttrAt,
	"attrs":   attrNoImplDataframe("attrs"),
	"axes":    attrNoImplDataframe("axes"),
	"columns": dataframeAttrColumns,
	"dtypes":  attrNoImplDataframe("dtypes"),
	"empty":   attrNoImplDataframe("empty"),
	"flags":   attrNoImplDataframe("flags"),
	"iat":     attrNoImplDataframe("iat"),
	"iloc":    attrNoImplDataframe("iloc"),
	"index":   dataframeAttrIndex,
	"loc":     attrNoImplDataframe("loc"),
	"ndim":    attrNoImplDataframe("ndim"),
	"shape":   dataframeAttrShape,
	"size":    attrNoImplDataframe("sizes"),
	"style":   attrNoImplDataframe("style"),
	"values":  attrNoImplDataframe("values"),
}

var dataframeMethods = map[string]*starlark.Builtin{
	"abs":               starlark.NewBuiltin("abs", methNoImpl("abs")),
	"add":               starlark.NewBuiltin("add", methNoImpl("add")),
	"add_prefix":        starlark.NewBuiltin("add_prefix", methNoImpl("add_prefix")),
	"add_suffix":        starlark.NewBuiltin("add_suffix", methNoImpl("add_suffix")),
	"agg":               starlark.NewBuiltin("agg", methNoImpl("agg")),
	"aggregate":         starlark.NewBuiltin("aggregate", methNoImpl("aggregate")),
	"align":             starlark.NewBuiltin("align", methNoImpl("align")),
	"all":               starlark.NewBuiltin("all", methNoImpl("all")),
	"any":               starlark.NewBuiltin("any", methNoImpl("any")),
	"append":            starlark.NewBuiltin("append", dataframeAppend),
	"apply":             starlark.NewBuiltin("apply", dataframeApply),
	"applymap":          starlark.NewBuiltin("applymap", methNoImpl("applymap")),
	"asfreq":            starlark.NewBuiltin("asfreq", methNoImpl("asfreq")),
	"asof":              starlark.NewBuiltin("asof", methNoImpl("asof")),
	"assign":            starlark.NewBuiltin("assign", methNoImpl("assign")),
	"astype":            starlark.NewBuiltin("astype", methNoImpl("astype")),
	"at_time":           starlark.NewBuiltin("at_time", methNoImpl("at_time")),
	"backfill":          starlark.NewBuiltin("backfill", methNoImpl("backfill")),
	"between_time":      starlark.NewBuiltin("between_time", methNoImpl("between_time")),
	"bfill":             starlark.NewBuiltin("bfill", methNoImpl("bfill")),
	"bool":              starlark.NewBuiltin("bool", methNoImpl("bool")),
	"boxplot":           starlark.NewBuiltin("boxplot", methMissing("boxplot")),
	"clip":              starlark.NewBuiltin("clip", methNoImpl("clip")),
	"combine":           starlark.NewBuiltin("combine", methNoImpl("combine")),
	"combine_first":     starlark.NewBuiltin("combine_first", methNoImpl("combine_first")),
	"compare":           starlark.NewBuiltin("compare", methNoImpl("compare")),
	"convert_dtypes":    starlark.NewBuiltin("convert_dtypes", methNoImpl("convert_dtypes")),
	"copy":              starlark.NewBuiltin("copy", methNoImpl("copy")),
	"corr":              starlark.NewBuiltin("corr", methNoImpl("corr")),
	"corrwith":          starlark.NewBuiltin("corrwith", methNoImpl("corrwith")),
	"count":             starlark.NewBuiltin("count", methNoImpl("count")),
	"cov":               starlark.NewBuiltin("cov", methNoImpl("cov")),
	"cummax":            starlark.NewBuiltin("cummax", methNoImpl("cummax")),
	"cummin":            starlark.NewBuiltin("cummin", methNoImpl("cummin")),
	"cumprod":           starlark.NewBuiltin("cumprod", methNoImpl("cumprod")),
	"cumsum":            starlark.NewBuiltin("cumsum", methNoImpl("cumsum")),
	"describe":          starlark.NewBuiltin("describe", methNoImpl("describe")),
	"diff":              starlark.NewBuiltin("diff", methNoImpl("diff")),
	"div":               starlark.NewBuiltin("div", methNoImpl("div")),
	"divide":            starlark.NewBuiltin("divide", methNoImpl("divide")),
	"dot":               starlark.NewBuiltin("dot", methNoImpl("dot")),
	"drop":              starlark.NewBuiltin("drop", dataframeDrop),
	"drop_duplicates":   starlark.NewBuiltin("drop_duplicates", dataframeDropDuplicates),
	"droplevel":         starlark.NewBuiltin("droplevel", methNoImpl("droplevel")),
	"dropna":            starlark.NewBuiltin("dropna", methNoImpl("dropna")),
	"duplicated":        starlark.NewBuiltin("duplicated", methNoImpl("duplicated")),
	"eq":                starlark.NewBuiltin("eq", methNoImpl("eq")),
	"equals":            starlark.NewBuiltin("equals", methNoImpl("equals")),
	"eval":              starlark.NewBuiltin("eval", methNoImpl("eval")),
	"ewm":               starlark.NewBuiltin("ewm", methNoImpl("ewm")),
	"expanding":         starlark.NewBuiltin("expanding", methNoImpl("expanding")),
	"explode":           starlark.NewBuiltin("explode", methNoImpl("explode")),
	"ffill":             starlark.NewBuiltin("ffill", methNoImpl("ffill")),
	"fillna":            starlark.NewBuiltin("fillna", methNoImpl("fillna")),
	"filter":            starlark.NewBuiltin("filter", methNoImpl("filter")),
	"first":             starlark.NewBuiltin("first", methNoImpl("first")),
	"first_valid_index": starlark.NewBuiltin("first_valid_index", methNoImpl("first_valid_index")),
	"floordiv":          starlark.NewBuiltin("floordiv", methNoImpl("floordiv")),
	"from_dict":         starlark.NewBuiltin("from_dict", methNoImpl("from_dict")),
	"from_records":      starlark.NewBuiltin("from_records", methNoImpl("from_records")),
	"ge":                starlark.NewBuiltin("ge", methNoImpl("ge")),
	"get":               starlark.NewBuiltin("get", methNoImpl("get")),
	"groupby":           starlark.NewBuiltin("groupby", dataframeGroupBy),
	"gt":                starlark.NewBuiltin("gt", methNoImpl("gt")),
	"head":              starlark.NewBuiltin("head", dataframeHead),
	"hist":              starlark.NewBuiltin("hist", methNoImpl("hist")),
	"idxmax":            starlark.NewBuiltin("idxmax", methNoImpl("idxmax")),
	"idxmin":            starlark.NewBuiltin("idxmin", methNoImpl("idxmin")),
	"infer_objects":     starlark.NewBuiltin("infer_objects", methNoImpl("infer_objects")),
	"info":              starlark.NewBuiltin("info", methNoImpl("info")),
	"insert":            starlark.NewBuiltin("insert", methNoImpl("insert")),
	"interpolate":       starlark.NewBuiltin("interpolate", methNoImpl("interpolate")),
	"isin":              starlark.NewBuiltin("isin", methNoImpl("isin")),
	"isna":              starlark.NewBuiltin("isna", methNoImpl("isna")),
	"isnull":            starlark.NewBuiltin("isnull", methNoImpl("isnull")),
	"items":             starlark.NewBuiltin("items", methNoImpl("items")),
	"iteritems":         starlark.NewBuiltin("iteritems", methNoImpl("iteritems")),
	"iterrows":          starlark.NewBuiltin("iterrows", methNoImpl("iterrows")),
	"itertuples":        starlark.NewBuiltin("itertuples", methNoImpl("itertuples")),
	"join":              starlark.NewBuiltin("join", methNoImpl("join")),
	"keys":              starlark.NewBuiltin("keys", methNoImpl("keys")),
	"kurt":              starlark.NewBuiltin("kurt", methNoImpl("kurt")),
	"kurtosis":          starlark.NewBuiltin("kurtosis", methNoImpl("kurtosis")),
	"last":              starlark.NewBuiltin("last", methNoImpl("last")),
	"last_valid_index":  starlark.NewBuiltin("last_valid_index", methNoImpl("last_valid_index")),
	"le":                starlark.NewBuiltin("le", methNoImpl("le")),
	"lookup":            starlark.NewBuiltin("lookup", methNoImpl("lookup")),
	"lt":                starlark.NewBuiltin("lt", methNoImpl("lt")),
	"mad":               starlark.NewBuiltin("mad", methNoImpl("mad")),
	"mask":              starlark.NewBuiltin("mask", methNoImpl("mask")),
	"max":               starlark.NewBuiltin("max", methNoImpl("max")),
	"mean":              starlark.NewBuiltin("mean", methNoImpl("mean")),
	"median":            starlark.NewBuiltin("median", methNoImpl("median")),
	"melt":              starlark.NewBuiltin("melt", methNoImpl("melt")),
	"memory_usage":      starlark.NewBuiltin("memory_usage", methNoImpl("memory_usage")),
	"merge":             starlark.NewBuiltin("merge", dataframeMerge),
	"min":               starlark.NewBuiltin("min", methNoImpl("min")),
	"mod":               starlark.NewBuiltin("mod", methNoImpl("mod")),
	"mode":              starlark.NewBuiltin("mode", methNoImpl("mode")),
	"mul":               starlark.NewBuiltin("mul", methNoImpl("mul")),
	"multiply":          starlark.NewBuiltin("multiply", methNoImpl("multiply")),
	"ne":                starlark.NewBuiltin("ne", methNoImpl("ne")),
	"nlargest":          starlark.NewBuiltin("nlargest", methNoImpl("nlargest")),
	"notna":             starlark.NewBuiltin("notna", methNoImpl("notna")),
	"notnull":           starlark.NewBuiltin("notnull", methNoImpl("notnull")),
	"nsmallest":         starlark.NewBuiltin("nsmallest", methNoImpl("nsmallest")),
	"nunique":           starlark.NewBuiltin("nunique", methNoImpl("nunique")),
	"pad":               starlark.NewBuiltin("pad", methNoImpl("pad")),
	"pct_change":        starlark.NewBuiltin("pct_change", methNoImpl("pct_change")),
	"pipe":              starlark.NewBuiltin("pipe", methNoImpl("pipe")),
	"pivot":             starlark.NewBuiltin("pivot", methNoImpl("pivot")),
	"pivot_table":       starlark.NewBuiltin("pivot_table", methNoImpl("pivot_table")),
	"plot":              starlark.NewBuiltin("plot", methMissing("plot")),
	"pop":               starlark.NewBuiltin("pop", methNoImpl("pop")),
	"pow":               starlark.NewBuiltin("pow", methNoImpl("pow")),
	"prod":              starlark.NewBuiltin("prod", methNoImpl("prod")),
	"product":           starlark.NewBuiltin("product", methNoImpl("product")),
	"quantile":          starlark.NewBuiltin("quantile", methNoImpl("quantile")),
	"query":             starlark.NewBuiltin("query", methNoImpl("query")),
	"radd":              starlark.NewBuiltin("radd", methNoImpl("radd")),
	"rank":              starlark.NewBuiltin("rank", methNoImpl("rank")),
	"rdiv":              starlark.NewBuiltin("rdiv", methNoImpl("rdiv")),
	"reindex":           starlark.NewBuiltin("reindex", methNoImpl("reindex")),
	"reindex_like":      starlark.NewBuiltin("reindex_like", methNoImpl("reindex_like")),
	"rename":            starlark.NewBuiltin("rename", methNoImpl("rename")),
	"rename_axis":       starlark.NewBuiltin("rename_axis", methNoImpl("rename_axis")),
	"reorder_levels":    starlark.NewBuiltin("reorder_levels", methNoImpl("reorder_levels")),
	"replace":           starlark.NewBuiltin("replace", methNoImpl("replace")),
	"resample":          starlark.NewBuiltin("resample", methNoImpl("resample")),
	"reset_index":       starlark.NewBuiltin("reset_index", dataframeResetIndex),
	"rfloordiv":         starlark.NewBuiltin("rfloordiv", methNoImpl("rfloordiv")),
	"rmod":              starlark.NewBuiltin("rmod", methNoImpl("rmod")),
	"rmul":              starlark.NewBuiltin("rmul", methNoImpl("rmul")),
	"rolling":           starlark.NewBuiltin("rolling", methNoImpl("rolling")),
	"round":             starlark.NewBuiltin("round", methNoImpl("round")),
	"rpow":              starlark.NewBuiltin("rpow", methNoImpl("rpow")),
	"rsub":              starlark.NewBuiltin("rsub", methNoImpl("rsub")),
	"rtruediv":          starlark.NewBuiltin("rtruediv", methNoImpl("rtruediv")),
	"sample":            starlark.NewBuiltin("sample", methNoImpl("sample")),
	"select_dtypes":     starlark.NewBuiltin("select_dtypes", methNoImpl("select_dtypes")),
	"sem":               starlark.NewBuiltin("sem", methNoImpl("sem")),
	"set_axis":          starlark.NewBuiltin("set_axis", methNoImpl("set_axis")),
	"set_flags":         starlark.NewBuiltin("set_flags", methNoImpl("set_flags")),
	"set_index":         starlark.NewBuiltin("set_index", methNoImpl("set_index")),
	"shift":             starlark.NewBuiltin("shift", methNoImpl("shift")),
	"skew":              starlark.NewBuiltin("skew", methNoImpl("skew")),
	"slice_shift":       starlark.NewBuiltin("slice_shift", methNoImpl("slice_shift")),
	"sort_index":        starlark.NewBuiltin("sort_index", methNoImpl("sort_index")),
	"sort_values":       starlark.NewBuiltin("sort_values", dataframeSortValues),
	"sparse":            starlark.NewBuiltin("sparse", methNoImpl("sparse")),
	"squeeze":           starlark.NewBuiltin("squeeze", methNoImpl("squeeze")),
	"stack":             starlark.NewBuiltin("stack", methNoImpl("stack")),
	"std":               starlark.NewBuiltin("std", methNoImpl("std")),
	"sub":               starlark.NewBuiltin("sub", methNoImpl("sub")),
	"subtract":          starlark.NewBuiltin("subtract", methNoImpl("subtract")),
	"sum":               starlark.NewBuiltin("sum", methNoImpl("sum")),
	"swapaxes":          starlark.NewBuiltin("swapaxes", methNoImpl("swapaxes")),
	"swaplevel":         starlark.NewBuiltin("swaplevel", methNoImpl("swaplevel")),
	"tail":              starlark.NewBuiltin("tail", methNoImpl("tail")),
	"take":              starlark.NewBuiltin("take", methNoImpl("take")),
	"to_clipboard":      starlark.NewBuiltin("to_clipboard", methMissing("to_clipboard")),
	"to_csv":            starlark.NewBuiltin("to_csv", methNoImpl("to_csv")),
	"to_dict":           starlark.NewBuiltin("to_dict", methNoImpl("to_dict")),
	"to_excel":          starlark.NewBuiltin("to_excel", methMissing("to_excel")),
	"to_feather":        starlark.NewBuiltin("to_feather", methMissing("to_feather")),
	"to_gbq":            starlark.NewBuiltin("to_gbq", methNoImpl("to_gbq")),
	"to_hdf":            starlark.NewBuiltin("to_hdf", methNoImpl("to_hdf")),
	"to_html":           starlark.NewBuiltin("to_html", methMissing("to_html")),
	"to_json":           starlark.NewBuiltin("to_json", methNoImpl("to_json")),
	"to_latex":          starlark.NewBuiltin("to_latex", methNoImpl("to_latex")),
	"to_markdown":       starlark.NewBuiltin("to_markdown", methNoImpl("to_markdown")),
	"to_numpy":          starlark.NewBuiltin("to_numpy", methNoImpl("to_numpy")),
	"to_parquet":        starlark.NewBuiltin("to_parquet", methNoImpl("to_parquet")),
	"to_period":         starlark.NewBuiltin("to_period", methNoImpl("to_period")),
	"to_pickle":         starlark.NewBuiltin("to_pickle", methMissing("to_pickle")),
	"to_records":        starlark.NewBuiltin("to_records", methNoImpl("to_records")),
	"to_sql":            starlark.NewBuiltin("to_sql", methNoImpl("to_sql")),
	"to_stata":          starlark.NewBuiltin("to_stata", methMissing("to_stata")),
	"to_string":         starlark.NewBuiltin("to_string", methNoImpl("to_string")),
	"to_timestamp":      starlark.NewBuiltin("to_timestamp", methNoImpl("to_timestamp")),
	"to_xarray":         starlark.NewBuiltin("to_xarray", methNoImpl("to_xarray")),
	"to_xml":            starlark.NewBuiltin("to_xml", methMissing("to_xml")),
	"transform":         starlark.NewBuiltin("transform", methNoImpl("transform")),
	"transpose":         starlark.NewBuiltin("transpose", methNoImpl("transpose")),
	"truediv":           starlark.NewBuiltin("truediv", methNoImpl("truediv")),
	"truncate":          starlark.NewBuiltin("truncate", methNoImpl("truncate")),
	"tshift":            starlark.NewBuiltin("tshift", methNoImpl("tshift")),
	"tz_convert":        starlark.NewBuiltin("tz_convert", methNoImpl("tz_convert")),
	"tz_localize":       starlark.NewBuiltin("tz_localize", methNoImpl("tz_localize")),
	"unstack":           starlark.NewBuiltin("unstack", methNoImpl("unstack")),
	"update":            starlark.NewBuiltin("update", methNoImpl("update")),
	"value_counts":      starlark.NewBuiltin("value_counts", methNoImpl("value_counts")),
	"var":               starlark.NewBuiltin("var", methNoImpl("var")),
	"where":             starlark.NewBuiltin("where", methNoImpl("where")),
	"xs":                starlark.NewBuiltin("xs", methNoImpl("xs")),
}

type starlarkMethod func(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)

func attrNoImplDataframe(attrName string) dataframeAttrImpl {
	return func(*DataFrame) (starlark.Value, error) {
		return nil, fmt.Errorf("dataframe.%s is not implemented. If you need this functionality to exist, file an issue at 'https://github.com/qri-io/starlib/issues' with the title 'dataframe.%s needs implementation'. Please first search if an issue exists already", attrName, attrName)
	}
}

func attrNoImplSeries(attrName string) seriesAttrImpl {
	return func(*Series) (starlark.Value, error) {
		return nil, fmt.Errorf("series.%s is not implemented. If you need this functionality to exist, file an issue at 'https://github.com/qri-io/starlib/issues' with the title 'series.%s needs implementation'. Please first search if an issue exists already", attrName, attrName)
	}
}

func methNoImpl(methodName string) starlarkMethod {
	return func(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
		return nil, fmt.Errorf("dataframe.%s is not implemented. If you need this functionality to exist, file an issue at 'https://github.com/qri-io/starlib/issues' with the title 'dataframe.%s needs implementation'. Please first search if an issue exists already", methodName, methodName)
	}
}

func methMissing(methodName string) starlarkMethod {
	return func(_ *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
		return nil, fmt.Errorf("dataframe.%s does not exist in this environment", methodName)
	}
}
