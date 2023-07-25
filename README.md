# lavasrc-plugin

[DisGoLink](https://github.com/disgoorg/disgolink) plugin for [LavaSrc](https://github.com/topi314/LavaSrc) for more track & playlist information.

## Installation

```shell
$ go get github.com/disgoorg/lavasrc-plugin
```

## Usage

```go
var track lavalink.Track
var lavasrcInfo lavasrc.TrackInfo
track.PluginInfo.Unmarshal(&lavasrcInfo)
```

```go
var playlist lavalink.Playlist
var lavasrcInfo lavasrc.PlaylistInfo
playlist.PluginInfo.Unmarshal(&lavasrcInfo)
```

## Troubleshooting

For help feel free to open an issue or reach out on [Discord](https://discord.gg/TewhTfDpvW)

## Contributing

Contributions are welcomed but for bigger changes we recommend first reaching out via [Discord](https://discord.gg/TewhTfDpvW) or create an issue to discuss your problems, intentions and ideas.

## License

Distributed under the [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/disgoorg/lavasrc-plugin/blob/master/LICENSE). See LICENSE for more information.
