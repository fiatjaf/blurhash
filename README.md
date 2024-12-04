blurhash
--------

simple blurhash encoder/decoder based on https://github.com/bbrks/go-blurhash.

## installation

```
go install github.com/fiatjaf/blurhash
```

## usage

```shell
~> blurhash encode -a -i ./assets/input.png
rAG+UJS[N4G?ZGD-;mWBt-XSIlVa-Et3R5xCS5~pi%9HTWj:Aa$OEWBRCies:pZNZm:x?E*R*D%tQ-:MyWBozIUs:j?V=XM%Nvo9ZKN%MwJs:
```

```shell
~> blurhash decode --hash='rAG+UJS[N4G?ZGD-;mWBt-XSIlVa-Et3R5xCS5~pi%9HTWj:Aa$OEWBRCies:pZNZm:x?E*R*D%tQ-:MyWBozIUs:j?V=XM%Nvo9ZKN%MwJs:' -o ./assets/output.png
```
