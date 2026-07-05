// 编码（保留字符转为%形式）
function encodeURIComponentSafe(str) {
    return encodeURIComponent(str).replace(/'/g, "%27"); // 显式处理单引号
}

// 解码
function decodeURIComponentSafe(encoded) {
    return decodeURIComponent(encoded.replace(/%27/g, "'")); // 还原单引号
}

// 异常捕获
function tryCatch(func){
    try{
        if(func){
            func();
        }
    }catch(e){
        Window.this.modal(<error caption="温馨提示">
                        <p>异常捕获</p>
                        <p>{e.message}</p>
                        <p>{e.stack}</p>
                    </error>);
    }
}

// 异常捕获异步执行
function tryCatchAsync(func) {
    setTimeout(function () {
        try {
            if (func) {
                func()
            }
        } catch (e) {
            Window.this.modal(<error caption="温馨提示">
                <p>异常捕获</p>
                <p>{e.message}</p>
                <p>{e.stack}</p>
            </error>);
        }
    }, 10);
}

function isStrEmpty(input){
    return input==undefined || (input=="")
}

function showInfo(message){
    Window.this.modal(<info caption="温馨提示">
                <p>{JSON.stringify(message,null,2)}</p>
            </info>);
}
/** 
    option：Window.this.selectFile(option)参数
        mode : "save"|"open"|"open-multiple"
        filter : "title|ext1;ext2", "HTML File (*.htm,*.html)|*.html;*.htm|All Files (*.*)|*.*"
        extension : default file extension, "html"
        caption : title of dialog, "Save As"
        path : initial directory  
    inputSelector：选择结果输出到控件选择器
 */
function selectFile(option,inputSelector){
    tryCatch(()=>{
        let r = Window.this.selectFile(option);
        if(r==null){
            return;
        }
        r = decodeURIComponentSafe(r);
        let element = document.$(inputSelector);
        if(element!=undefined){
            element.value = r.substring(7);
        }
    });
}