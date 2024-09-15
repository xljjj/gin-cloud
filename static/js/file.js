// 查看文件
function viewFile(fileId) {

}

// 下载文件
function downloadFile(fileId) {
    let curURL=window.location.href
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
    let curURL=window.location.href
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