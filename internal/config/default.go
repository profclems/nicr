package config

const defaultConfig = `
{
	"unknown_files_folder" : "Other",
	"known_files": [
		{
			"extensions": [
				"jpg",
				"jpeg",
				"png",
				"gif",
				"webp",
				"cr2",
				"tif",
				"bmp",
				"heif",
				"jxr",
				"psd",
				"ico",
				"dwg",
				"svg"
			],
			"folder": "Pictures",
			"exempt_files": false
		},
		{
			"extensions": [
				"mp4",
				"m4v",
				"mkv",
				"webm",
				"mov",
				"avi",
				"wmv",
				"mpg",
				"flv",
				"3gp"
			],
			"folder": "Videos",
			"exempt_files": false
		},
		{
			"extensions": [
				"wasm",
				"dex",
				"dey",
				"exe",
				"dmg",
				"rpm",
				"deb",
				"pkg",
				"apk"
			],
			"folder": "Applications",
			"exempt_files": false
		},
		{
			"extensions": [
				"woff",
				"woff2",
				"ttf",
				"otf"
			],
			"folder": "Fonts",
			"exempt_files": false
		},
		{
			"extensions": [
				"doc",
				"docx",
				"xls",
				"xlsx",
				"ppt",
				"pptx",
				"pdf",
				"epub",
				"rtf",
				"txt"
			],
			"folder": "Documents",
			"exempt_files": false
		},
		{
			"extensions": [
				"mid",
				"mp3",
				"m4a",
				"ogg",
				"flac",
				"wav",
				"amr",
				"aac"
			],
			"folder": "Audio",
			"exempt_files": false
		},
		{
			"extensions": [
				"zip",
				"tar",
				"rar",
				"gz",
				"bz2",
				"7z",
				"xz",
				"zstd",
				"swf",
				"iso",
				"eot",
				"ps",
				"nes",
				"crx",
				"cab",
				"ar",
				"Z",
				"lz",
				"elf",
				"dcm"
			],
			"folder": "Archive",
			"exempt_files": false
		},
		{
			"extensions": [
				"sqlite",
				"sql"
			],
			"folder": "Database",
			"exempt_files": false
		},
		{
			"extensions": [
				"download",
				"crdownload",
				".DS_Store"
			],
			"folder": "",
			"exempt_files": true
		}
	]
}
`
