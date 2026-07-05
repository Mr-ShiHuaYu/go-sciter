// 统一处理按钮事件
let modName = "";
function setModName(name) {
    modName = name;
}
document.on("click", "select#demolist > option", function (evt, element) {
    tryCatch(() => {
        let src = element.attributes["src"];
        let id = element.attributes["value"];
        let url = src;
        if (url == "" || url == undefined) {
            url = id;
        }
        url = modName + "/" + url + ".html"
        // document.$("#sampleFrame").src = url;
        document.$("frame#sampleFrame").frame.loadFile(url);
    });
});