const buybacks = {
    setUp: () => {
        buybacks.formatFloatFields()
    },

    formatFloatFields: () => {
        $(function() {
            $('#Buybacks').find('.float-number').each(function () {
                var originalNumber = parseFloat($(this).text());
                var formattedNumber = originalNumber.toFixed(7);
                $(this).text(formattedNumber);
            });
            $('#Buybacks').find('.float-number-two').each(function () {
                var originalNumber = parseFloat($(this).text());
                var formattedNumber = originalNumber.toFixed(2);
                $(this).text(formattedNumber);
            });
        });
    },
}

export default buybacks;