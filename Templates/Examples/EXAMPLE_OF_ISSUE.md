Greeting flameshot maintainers and users,

Sorry for the bother, but I have an issue that I cannot figure out myself.

## Describe the Issue

Flameshot leaves the screen to be taken as floating in a black background(zoom out), while the selection region stays at the normal size when taking screenshots.

### Flameshot Version

```
$ flameshot --version
Flameshot v12.1.0 (-)
Compiled with Qt 5.15.14
```

### Operating System

```
$ uname -r
6.11.2-4-MANJARO
```

### What Happended

The following picture taken by Spectacle describes the issue.

![](https://private-user-images.githubusercontent.com/62749885/333991312-1535a90e-fb2b-45d8-835f-78d3e024da4e.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MzEzODMyMTksIm5iZiI6MTczMTM4MjkxOSwicGF0aCI6Ii82Mjc0OTg4NS8zMzM5OTEzMTItMTUzNWE5MGUtZmIyYi00NWQ4LTgzNWYtNzhkM2UwMjRkYTRlLnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNDExMTIlMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjQxMTEyVDAzNDE1OVomWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPTBmMmFjYzRjNzU4M2Q4ZDI5ZjJiZWRjZjJmNTZjZDUxNjhhZDZjYTIzMGIxOGE0MDljOWEzOTczOWJjNjFkMjQmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0In0.ezbRrRS7TKoslLXOWdPsP45W-hLYdjbC9F10EQopteI)

And the output of `flameshot gui` in konsole.(Not for sure, those errors will be taken in count.)

```
QFont::fromString: Invalid description 'Hack Nerd Font,12,-1,5,900,0,0,0,0,0,0,0,0,0,0,1,Regular'
QFont::fromString: Invalid description 'Hack Nerd Font,12,-1,5,900,0,0,0,0,0,0,0,0,0,0,1,Regular'
QFont::fromString: Invalid description 'Hack Nerd Font,12,-1,5,900,0,0,0,0,0,0,0,0,0,0,1,Regular'
QFont::fromString: Invalid description 'Hack Nerd Font,12,-1,5,900,0,0,0,0,0,0,0,0,0,0,1,Regular'
qt.qpa.wayland: Wayland does not support QWindow::requestActivate()
qt.qpa.wayland: Wayland does not support QWindow::requestActivate()
qt.qpa.wayland: Wayland does not support QWindow::requestActivate()
flameshot: info: Screenshot aborted.
```

Also here is a video if the picture does not describe the issue clearly.

<video src="https://private-user-images.githubusercontent.com/62749885/331951334-a27e6af8-34c9-416c-9eed-94be7b768fb1.mp4?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MzEzOTg4ODgsIm5iZiI6MTczMTM5ODU4OCwicGF0aCI6Ii82Mjc0OTg4NS8zMzE5NTEzMzQtYTI3ZTZhZjgtMzRjOS00MTZjLTllZWQtOTRiZTdiNzY4ZmIxLm1wND9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNDExMTIlMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjQxMTEyVDA4MDMwOFomWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPTA4ZDljYjZhMDhkYzA0MTg5ZDhjMGY1MzEyMTQzOGVlZThmYmJlMmZhNzhhMDFmMmExODJhYWRmM2FkMmVlMjYmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0In0.-iQq_IHk50wnkh5o6h6pJ-oWq9mpTl8cd-TSOlaqzXs" controls></video>

### When did It Happen

Run `flameshot gui` or click the take screenshot button in GUI.

## How to Reproduce

1. Set the display scaling to 125% or 150%.
2. Run `flameshot gui` or click the take screenshot button in GUI.

## Things I have Tried

1. Use SEO search it, but not much related content, if I not missing something.
2. Follow the [troubleshooting](https://flameshot.org/docs/guide/troubleshooting/#i-have-fractional-scaling-on-my-monitors-e-g-150) section of flameshot. Use command `env QT_AUTO_SCREEN_SCALE_FACTOR="1.25" QT_SCREEN_SCALE_FACTORS="" flameshot gui`, not work.
3. Use command `env XDG_SESSION_TYPE=x11 flameshot gui`, works, but keystrokes will never be read when Legacy X11 App Support is set to Never.
4. Use command `env QT_QPA_PLATFORM=xcb flameshot gui`, works, same as X11.
5. Set Kwin rules follow the instruction of [issue#3073](https://github.com/flameshot-org/flameshot/issues/3073#issuecomment-1740187784) , not work.
6. Set display scaling to 100%, works.
7. Use appimage version, not work.

## Additional Information

### Display Configuration of KDE

![](https://private-user-images.githubusercontent.com/62749885/334009403-c4d5a6c6-90e2-4d15-b601-7cea4576a80c.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MzEzOTg4ODgsIm5iZiI6MTczMTM5ODU4OCwicGF0aCI6Ii82Mjc0OTg4NS8zMzQwMDk0MDMtYzRkNWE2YzYtOTBlMi00ZDE1LWI2MDEtN2NlYTQ1NzZhODBjLnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNDExMTIlMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjQxMTEyVDA4MDMwOFomWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPWE2N2Q2NmFmN2VmYzc4MDk5N2ZiYzIzYTJkMWU0NmVkOWE2MzQxZTcyZDEzZTM2YjUxZDEwMzQ1NGRjMDQ0MTQmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0In0.-N4jMfNn-ZAiV6G_AkoftIpVxlYKvjBnG_3yNasBWf0)
### Hardware & Display Server

```
$ inxi -xxG
Graphics:
  Device-1: NVIDIA AD107M [GeForce RTX 4060 Max-Q / Mobile]
    vendor: Micro-Star MSI driver: nvidia v: 550.78 arch: Lovelace pcie:
    speed: 2.5 GT/s lanes: 2 ports: active: none empty: DP-1, DP-2, HDMI-A-1,
    eDP-1 bus-ID: 01:00.0 chip-ID: 10de:28a0
  Device-2: AMD Raphael vendor: Micro-Star MSI driver: amdgpu v: kernel
    arch: RDNA-2 pcie: speed: 16 GT/s lanes: 16 ports: active: eDP-2 empty: DP-3,
    DP-4, DP-5, Writeback-1 bus-ID: 06:00.0 chip-ID: 1002:164e temp: 46.0 C
  Device-3: Bison HD Webcam driver: uvcvideo type: USB rev: 2.0
    speed: 480 Mb/s lanes: 1 bus-ID: 5-1.1:3 chip-ID: 5986:211c
  Display: wayland server: X.org v: 1.21.1.13 with: Xwayland v: 23.2.6
    compositor: kwin_wayland driver: X: loaded: amdgpu,nvidia dri: radeonsi
    gpu: nvidia,amdgpu display-ID: 0
  Monitor-1: eDP-2 res: 2048x1152 size: N/A
  API: EGL v: 1.5 platforms: device: 0 drv: nvidia device: 2 drv: radeonsi
    device: 3 drv: swrast gbm: drv: nvidia surfaceless: drv: nvidia wayland:
    drv: radeonsi x11: drv: radeonsi inactive: device-1
  API: OpenGL v: 4.6.0 compat-v: 4.5 vendor: amd mesa v: 24.0.6-manjaro1.1
    glx-v: 1.4 direct-render: yes renderer: AMD Radeon Graphics (radeonsi
    raphael_mendocino LLVM 17.0.6 DRM 3.57 6.9.0-1-MANJARO)
    device-ID: 1002:164e display-ID: :1.0
  API: Vulkan v: 1.3.279 surfaces: xcb,xlib,wayland device: 0
    type: discrete-gpu driver: nvidia device-ID: 10de:28a0 device: 1
    type: integrated-gpu driver: mesa radv device-ID: 1002:164e
```

### DE Version

```
$ pacman -Qs plasma-desktop
local/plasma-desktop 6.0.4-1 (plasma)
    KDE Plasma Desktop
```


English is not my native lanuage; please excuse typing errors. And thanks in advance.

Best regards