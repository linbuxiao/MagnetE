# MagnetE

## Usage

```bash
# search magnet
magnete -q "大地惊雷"
```

```yaml
# output
    - title: 梦幻天堂·龙网(killman.net).720p.大地惊雷
      magnet: magnet:?xt=urn:btih:AEC9862F20C66D3F34FEDA71B3C5A56EE4230050&dn=%E6%A2%A6%E5%B9%BB%E5%A4%A9%E5%A0%82%C2%B7%E9%BE%99%E7%BD%91%28killman.net%29.720p.%E5%A4%A7%E5%9C%B0%E6%83%8A%E9%9B%B7&tr=http%3A%2F%2Ftracker.ktxp.com%3A6868%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A7070%2Fannounce&tr=udp%3A%2F%2Ftracker.ktxp.com%3A6868%2Fannounce&tr=udp%3A%2F%2Ftracker.ktxp.com%3A7070%2Fannounce&tr=http%3A%2F%2Fbtfans.3322.org%3A8000%2Fannounce&tr=http%3A%2F%2Fbtfans.3322.org%3A8080%2Fannounce&tr=http%3A%2F%2Fbtfans.3322.org%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.bittorrent.am%2Fannounce&tr=udp%3A%2F%2Ftracker.bitcomet.net%3A8080%2Fannounce&tr=http%3A%2F%2Ftk3.5qzone.net%3A8080%2F&tr=http%3A%2F%2Ftracker.btzero.net%3A8080%2Fannounce&tr=http%3A%2F%2Fscubt.wjl.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fbt.popgo.net%3A7456%2Fannounce&tr=http%3A%2F%2Fthetracker.org%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%2Fannounce&tr=http%3A%2F%2Ftracker.dmhy.org%3A8000%2Fannounce&tr=http%3A%2F%2Fbt.titapark.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.tjgame.enorth.com.cn%3A8000%2Fannounce&
      size: "3.1G"
    - title: 【大地惊雷】【高清BDRIP-RMVB.中字】【2010最新美国高分票房西部动作大片】
      magnet: magnet:?xt=urn:btih:3D03726F7C4E265184661A92FC83D2B40ECAA3D0&dn=%E3%80%90%E5%A4%A7%E5%9C%B0%E6%83%8A%E9%9B%B7%E3%80%91%E3%80%90%E9%AB%98%E6%B8%85BDRIP-RMVB.%E4%B8%AD%E5%AD%97%E3%80%91%E3%80%902010%E6%9C%80%E6%96%B0%E7%BE%8E%E5%9B%BD%E9%AB%98%E5%88%86%E7%A5%A8%E6%88%BF%E8%A5%BF%E9%83%A8%E5%8A%A8%E4%BD%9C%E5%A4%A7%E7%89%87%E3%80%91&tr=http%3A%2F%2Ftracker.ktxp.com%3A6868%2Fannounce&tr=http%3A%2F%2Ftracker.ktxp.com%3A7070%2Fannounce&tr=udp%3A%2F%2Ftracker.ktxp.com%3A6868%2Fannounce&tr=udp%3A%2F%2Ftracker.ktxp.com%3A7070%2Fannounce&tr=http%3A%2F%2Fbtfans.3322.org%3A8000%2Fannounce&tr=http%3A%2F%2Fbtfans.3322.org%3A8080%2Fannounce&tr=http%3A%2F%2Fbtfans.3322.org%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.bittorrent.am%2Fannounce&tr=udp%3A%2F%2Ftracker.bitcomet.net%3A8080%2Fannounce&tr=http%3A%2F%2Ftk3.5qzone.net%3A8080%2F&tr=http%3A%2F%2Ftracker.btzero.net%3A8080%2Fannounce&tr=http%3A%2F%2Fscubt.wjl.cn%3A8080%2Fannounce&tr=http%3A%2F%2Fbt.popgo.net%3A7456%2Fannounce&tr=http%3A%2F%2Fthetracker.org%2Fannounce&tr=http%3A%2F%2Ftracker.prq.to%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%2Fannounce&tr=http%3A%2F%2Ftracker.dmhy.org%3A8000%2Fannounce&tr=http%3A%2F%2Fbt.titapark.com%3A2710%2Fannounce&tr=http%3A%2F%2Ftracker.tjgame.enorth.com.cn%3A8000%2Fannounce&
      size: "4.6G"
      ...
```

## Develop Guide

1. Add rule in `rules.yaml`.

2. Generate a new client:
    ```bash
    make gen-collector
    ```

3. You get it.

