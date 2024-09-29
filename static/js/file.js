// 查看文件
function viewFile(fileId) {
    const statusDiv = document.getElementById('status');
    fetch(`/cloud/viewFile?fileId=${encodeURIComponent(fileId)}`, {
        method: 'GET',
    })
        .then(response => response.text())
        .then(filePath => {
            if (filePath){
                window.open(filePath, '_blank');
            }else{
                statusDiv.textContent='文件不存在';
            }
        })
        .catch(error => {
            statusDiv.textContent = '因网络原因文件查看失败';
        });
}

// 下载文件
function downloadFile(fileId) {
    const curURL=window.location.href
    // 创建一个隐藏的 <a> 标签用于下载
    const link = document.createElement('a');
    link.href = `/cloud/downloadFile?fileId=${encodeURIComponent(fileId)}`;
    link.download = ''; // 触发下载
    document.body.appendChild(link); // 将链接添加到 DOM
    link.click(); // 触发点击事件，开始下载
    document.body.removeChild(link); // 下载完成后移除链接
    window.location.href = curURL;
}

// 删除文件
function deleteFile(fileId) {
    const curURL=window.location.href
    const statusDiv = document.getElementById('status');
    if (confirm('确定删除此文件？')) {
        fetch(`/cloud/deleteFile?fileId=${encodeURIComponent(fileId)}`, {
            method: 'DELETE',
        })
            .then(response => response.json())
            .then(data => {
                if (data.error){
                    statusDiv.textContent='文件删除失败，原因：'+data.error;
                }else{
                    alert("文件删除成功！")
                    window.location.href = curURL;
                }
            })
            .catch(error => {
                statusDiv.textContent = '因网络原因文件删除失败';
            });
    }
}

// 查看文件夹
function viewFolder(folderId){
    window.location.href=`/cloud/file?fId=${encodeURIComponent(folderId)}`
}

// 新建文件夹
function createFolder(parentId){
    const statusDiv = document.getElementById('status');
    const curURL=window.location.href

    // 弹出输入框让用户输入文件夹名
    const folderName = prompt("请输入新文件夹的名称:");

    if (!folderName) {
        statusDiv.textContent = '未输入文件夹名！';
        return;
    }

    const data = {
        folderName: folderName,
        parentId: parentId
    };

    fetch('/cloud/addFolder', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    }).then(response => response.json())
        .then(data => {
            if (data.error){
                statusDiv.textContent='文件夹创建失败，原因：'+data.error;
            }else{
                alert("文件夹创建成功！")
                window.location.href = curURL;
            }
        })
        .catch(error => {
            statusDiv.textContent = '因网络原因文件夹创建失败';
        });
}

// 修改文件夹
function updateFolder(parentId,folderId){
    const statusDiv = document.getElementById('status');
    const curURL=window.location.href

    // 弹出输入框让用户输入文件夹名
    const folderName = prompt("请输入新文件夹的名称:");

    if (!folderName) {
        statusDiv.textContent = '未输入文件夹名！';
        return;
    }

    const data = {
        folderName: folderName,
        folderId: folderId,
        parentId: parentId
    };

    fetch('/cloud/updateFolder', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    }).then(response => response.json())
        .then(data => {
            if (data.error){
                statusDiv.textContent='文件夹更名失败，原因：'+data.error;
            }else{
                alert("文件夹更名成功！")
                window.location.href = curURL;
            }
        })
        .catch(error => {
            statusDiv.textContent = '因网络原因文件夹更名失败';
        });
}

// 删除文件夹
function deleteFolder(folderId) {
    const curURL=window.location.href
    const statusDiv = document.getElementById('status');
    if (confirm('确定删除此文件夹？')) {
        fetch(`/cloud/deleteFolder?fId=${encodeURIComponent(folderId)}`, {
            method: 'DELETE',
        })
            .then(response => response.json())
            .then(data => {
                if (data.error){
                    statusDiv.textContent='文件删除失败，原因：'+data.error;
                }else{
                    alert("文件夹删除成功！")
                    window.location.href = curURL;
                }
            })
            .catch(error => {
                statusDiv.textContent = '因网络原因文件夹删除失败';
            });
    }
}

// 浏览器回退
function goBack() {
    window.history.back();
}