const buybacks = {
    setUp: () => {
        buybacks.formatFloatFields()
    },

    formatFloatFields: () => {
        $(function() {
            $('#Buybacks').find('.float-number').each(function () {
                var originalNumber = parseFloat($(this).text());
                var formattedNumber = originalNumber.toFixed(4);
                $(this).text(formattedNumber);
            });
        });
    },
}

export default buybacks;