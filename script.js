var options = {
    "key": "rzp_test_2V4Q2lxFE0LPwf",
    "currency": "INR",
    "name": "Lightsail Software Services",
    "description": "Test Transaction",
    "image": "https://razorpay.com/build/browser/static/razorpay-logo.5cdb58df.svg",
    "handler": function (response){
        document.getElementById("payment-id").innerHTML = response.razorpay_payment_id;
        alert(response.razorpay_payment_id);
        alert(response.razorpay_order_id);
        alert(response.razorpay_signature)
    },
    "prefill": {
        "name": "Mehul Sharma",
        "email": "mehulsharma@gmail.com",
        "contact": "9999999999"
    },
    "notes": {
        "address": "BLR"
    },
    "theme": {
        "color": "#6f1632"
    }
};
let orderId;
document.getElementById('rzp-create-order').onclick = function(e){
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        var response;
        if (this.readyState == 4 && this.status == 200) {
            response = JSON.parse(this.responseText)
            console.log(response)
            orderId = response["id"]
            document.getElementById("order_id").setAttribute('value', orderId)
        }
    };
    const amount = document.getElementById("amount").value
    xhttp.open("GET", `http://localhost:8081/order?amount=${amount}`, true);
    xhttp.send();

}
document.getElementById('rzp-confirm-order-id').onclick = function(e){
    const order_id = document.getElementById("order_id").value;
    alert(order_id)
    options["order_id"] = String(order_id)
}
document.getElementById('rzp-pay').onclick = function(e){
    var rzp1 = new Razorpay(options);
    rzp1.on('payment.failed', function (response){
        alert(response.error.code);
        alert(response.error.description);
        alert(response.error.source);
        alert(response.error.step);
        alert(response.error.reason);
        alert(response.error.metadata.order_id);
        alert(response.error.metadata.payment_id);
    });
    rzp1.open();
    e.preventDefault();
}