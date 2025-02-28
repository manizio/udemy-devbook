$("#nova-publicacao").on("submit", createPost)
$("#edit-post").on("click", updatePost)
$(".delete-post").on("click", deletePost)


$(document).on('click', '.like-post', likePost)
$(document).on('click', '.unlike-post', unlikePost)

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

function likePost(event) {

    const clicked = $(event.target)
    const postID = clicked.closest("div").data('post-id')

    clicked.prop('disabled', true)
    $.ajax({
        url: `/posts/${postID}/like`,
        method: "POST"
    }).done(function() {
        const likeCounter = clicked.next("span")
        const likes = parseInt(likeCounter.text())

        likeCounter.text(likes + 1)
        clicked.addClass('unlike-post')
        clicked.addClass('text-danger')
        clicked.removeClass('like-post')
    }).fail(function() {
        alert("Erro ao curtir")
    }).always(function() {
        clicked.prop('disabled', false)
    })
}

function unlikePost(event) {
    event.preventDefault()

    const clicked = $(event.target)
    const postID = clicked.closest('div').data('post-id')

    clicked.prop('disabled', true)
    $.ajax({
        url: `/posts/${postID}/unlike`,
        method: "POST"
    }).done(function() {
        const likeCounter = clicked.next("span")
        const likes = parseInt(likeCounter.text())

        likeCounter.text(likes - 1)

        clicked.removeClass('unlike-post')
        clicked.removeClass('text-danger')
        clicked.addClass('like-post')

    }).always(function() {
        clicked.prop('disabled', false)
    })
}

function updatePost(){ 
    $(this).prop('disabled', true)

    const postID = $(this).data('post-id')
    $.ajax({
        url: `/posts/${postID}`,
        method: "PUT",
        data: {
            title: $("#title").val(),
            content: $("#content").val()
        }
    }).done(function() {
        alert("Editado")
    }).fail(function() {
        alert("erro ao editar publicação")
    }).always(function() {
        $("#edit-post").prop('disabled', false)
    })

}

function deletePost(event){
    event.preventDefault()

    const clicked = $(event.target)
    const post = clicked.closest("div")
    const postID = post.data("post-id")

    clicked.prop('disabled', true)

    $.ajax({
        url: `/posts/${postID}`,
        method: "DELETE",
    }).done(function (){
        post.fadeOut("slow", function() {
            $(this).remove()
        })
    }).fail(function() {
        alert("erro ao excluir publicação")
    })
}
