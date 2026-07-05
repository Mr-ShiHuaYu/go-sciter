// 拖动文件属性
class FileDropZone extends Element {
    files = []; // filtered files

    ondragaccept(evt) {
        if (evt.detail.dataType == "file") {
            this.files = evt.detail.data;
            if (!Array.isArray(this.files))
                this.files = [this.files];
            return true; // accept only files
        }
    }

    ondragenter(evt) {
        this.classList.add("active-target");
        return true;
    }

    ondragleave(evt) {
        this.classList.remove("active-target");
        return true;
    }

    ondrag(evt) {
        console.log(evt.x, evt.y);
        return true;
    }

    ondrop(evt) {
        this.classList.remove("active-target");
        // document.$("#ffmpegPath").value = this.files[0];
        if (this.files != undefined && this.files.length > 0) {
            let fileName = this.files[0];
            if (fileName != undefined && fileName != null) {
                const realFileName = fileName.substring(7);
                this.value = decodeURIComponent(realFileName.replace(/%27/g, "'"));
                // 创建自定义事件对象
                var event = new Event("change", {
                    bubbles: true,  // 允许事件冒泡
                    cancelable: true
                });
                event.data = { detail: realFileName }
                this.dispatchEvent(event);
            }
        }
        return true;
    }

}