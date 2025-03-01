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
        Swal.fire("Ops...", "Erro ao criar publicação", "error")
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
        Swal.fire("Ops...", "Erro ao curtir", "error")
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

function updatePost() {
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
        Swal.fire(
            'Sucesso!',
            'Publicação Atualizada com Sucesso',
            'success'
        ).then(function() {
            window.location = "/home"
        })
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao atualizar a publicação", "error")
    }).always(function() {
        $("#edit-post").prop('disabled', false)
    })

}

function deletePost(event) {
    event.preventDefault()

    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja excluir essa publicação?",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirm){
        if (!confirm.value) return;
        const clicked = $(event.target)
        const post = clicked.closest("div")
        const postID = post.data("post-id")

        clicked.prop('disabled', true)

        $.ajax({
            url: `/posts/${postID}`,
            method: "DELETE",
        }).done(function() {
            post.fadeOut("slow", function() {
                $(this).remove()
                })
        }).fail(function() {
            Swal.fire("Ops...", "Erro ao excluir publicação", "error")

        })
    })
}
