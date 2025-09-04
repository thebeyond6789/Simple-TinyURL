document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('shortenForm');
    const originalUrlInput = document.getElementById('originalUrl');
    const resultDiv = document.getElementById('result');
    const shortenedUrlElement = document.getElementById('shortenedUrl');
    const errorDiv = document.getElementById('error');
    const errorMessageElement = document.getElementById('errorMessage');

    form.addEventListener('submit', async (e) => {
        e.preventDefault();
        
        const originalUrl = originalUrlInput.value.trim();
        if (!originalUrl) {
            showError('Please enter a valid URL');
            return;
        }

        try {
            const response = await fetch('/api/shorten', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ url: originalUrl })
            });

            const data = await response.json();

            if (response.ok) {
                showResult(`${window.location.origin}/${data.slug}`);
            } else {
                showError(data.error || 'An error occurred');
            }
        } catch (error) {
            showError('Failed to connect to the server');
        }
    });

    function showResult(shortenedUrl) {
        shortenedUrlElement.href = shortenedUrl;
        shortenedUrlElement.textContent = shortenedUrl;
        resultDiv.classList.remove('hidden');
        errorDiv.classList.add('hidden');
    }

    function showError(message) {
        errorMessageElement.textContent = message;
        errorDiv.classList.remove('hidden');
        resultDiv.classList.add('hidden');
    }
});