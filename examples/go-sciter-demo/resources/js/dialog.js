export function ImgShowWindow(props) {
    const src = props.src;
    return <html window-frame="default"
        window-height="600" window-width="800">
        <head><title>图片展示</title></head>
        <body style="background-color:white;vertical-align:middle;horizontal-align:center;">
            <img src={src} style="width:700;" />
        </body>
    </html>;
}