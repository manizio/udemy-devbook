$("#unfollow").on("click", unfollow)
$("#follow").on("click", follow)

function unfollow() {
    const userID = $(this).data("user-id")
    $(this).prop('disabled', true)

    $.ajax({
        url: `/users/${userID}/unfollow`,
        method: "POST"
    }).done(function() {
        window.location = `/users/${userID}`
    }).fail(function () {
        Swal.fire("Ops...", "Erro ao parar de seguir o usuário", "error")
        $('#unfollow').prop('disabled', false)
    })
}

function follow() {
    const userID = $(this).data('user-id')
    $(this).prop('disabled', true)

    $.ajax({
        url: `/users/${userID}/follow`,
        method: "POST"
    }).done(function() {
        window.location = `/users/${userID}`
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao seguir o usuário", "error")
        $('#follow').prop('disabled', false)
    })

}
