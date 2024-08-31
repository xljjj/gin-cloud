document.addEventListener('DOMContentLoaded', function () {
    const deleteForms = document.querySelectorAll('.delete-section form');
    deleteForms.forEach(form => {
        form.addEventListener('submit', function (event) {
            if (!confirm('确认要删除该用户吗？')) {
                event.preventDefault(); // 防止表单提交
            }
        });
    });
});