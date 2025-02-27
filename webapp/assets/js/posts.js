$("#nova-publicacao").on("submit", createPost)
function createPost(event) {
    event.preventDefault()

    $.ajax({
        url: "/posts",
        method: "POST",
        data: {
            title: $("#title").val(),
            content: $("#content").val(),
        }
    }).done(function() {
        window.location = "/home"
    }).fail(function() {
        alert("Erro ao criar a publicação")
    })
}
