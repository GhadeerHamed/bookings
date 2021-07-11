function buildJSHandler(csrf, id) {
    document.getElementById("check-availability-btn").addEventListener("click", function () {
        let html = `
        <form id="check-availability-form" action="" method="post" novalidate class="needs-validation">
            <div class="row">
                <div class="col">
                    <div class="row" id="reservation-dates-modal">
                        <div class="col">
                            <input disabled autocomplete="off" required class="form-control" name="start" id="start" placeholder="Arrival" type="text" >
                        </div>
                        <div class="col">
                            <input disabled autocomplete="off" required class="form-control" name="end" id="end" placeholder="Departure" type="text" >
                        </div>
                    </div>
                </div>
            </div>
        </form>
    `


        attention().custom({
            msg: html,
            title: "Choose your dates",
            willOpen: (toast) => {
                const elem = document.getElementById("reservation-dates-modal")
                const rp = new DateRangePicker(elem, {
                    format: "yyyy-mm-dd",
                    showOnFocus: true,
                    minDate: new Date()
                });
            },

            didOpen: () => {
                document.getElementById("start").removeAttribute('disabled')
                document.getElementById("end").removeAttribute('disabled')
            },

            callback: function (r) {
                let form = document.getElementById("check-availability-form")
                let formData = new FormData(form)
                formData.append('csrf_token', csrf)
                formData.append('room_id', id)

                fetch("/search-availability-json", {
                    method: "post",
                    body: formData
                })
                    .then(response => response.json())
                    .then(data => {
                        if (data.ok) {
                            attention().custom({
                                icon: "success",
                                showConfirmButton: false,
                                msg: '<p>Room is available</p>'
                                    + '<p><a href="/book-room?id=' + data.room_id
                                    + '&s=' + data.start_date
                                    + '&e=' + data.end_date
                                    + '" class="btn btn-primary">'
                                    + 'Book Now!</a></p>'
                            })
                        } else {
                            attention().error({
                                msg: "No availability"
                            })
                        }
                    })
            }
        }).then(undefined)
    })
}