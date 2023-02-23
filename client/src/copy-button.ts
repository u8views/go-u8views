export function initCopyCodeButtons() {
    const $copyCodeButtons = document.querySelectorAll<HTMLElement>(".js-copy-code-button");

    $copyCodeButtons.forEach(($button) => {
        $button.addEventListener("click", function (event) {
            const $copyCheck = (event.target as Element).querySelector<HTMLElement>(".js-copy-code-check");
            const $copyDone = (event.target as Element).querySelector<HTMLElement>(".js-copy-code-done");

            $copyCheck.style.display = "none";
            $copyDone.style.display = "block";
            $button.style.animationName = "github-button";

            codeToClipboard($button);

            setTimeout(() => {
                $copyCheck.style.display = "block";
                $copyDone.style.display = "none";
                $button.style.animationName = "none";

            }, 2000);
        });
    });
}

function codeToClipboard($button) {
    const code = $button.parentElement.querySelector(".js-code-for-copy").innerText.trim();

    navigator.clipboard.writeText(code).catch(console.error);
}
