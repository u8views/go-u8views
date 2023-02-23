export function initCopyCodeButtons(){
    const $copyCodeButtons = document.querySelectorAll<HTMLElement>(".js-copy-code-button");

    if ($copyCodeButtons) {
        $copyCodeButtons.forEach((button) => {
            button.addEventListener("click", (e) => {

                const $copyCheck = (e.target as Element).querySelector<HTMLElement>('.js-copy-code-check')
                const $copyDone = (e.target as Element).querySelector<HTMLElement>('.js-copy-code-done')

                $copyCheck.style.display = 'none'
                $copyDone.style.display = 'block'
                button.style.animationName = "github-button"

                copyToClipboard(button)

                setTimeout(() => {
                    $copyCheck.style.display = 'block'
                    $copyDone.style.display = 'none'
                    button.style.animationName = "none"

                }, 2000);
            });
        });
    }
}

function copyToClipboard(button) {
    const $codeForCopy = (button.parentElement.querySelector('.js-code-for-copy') as HTMLParagraphElement)
        .textContent.trim();

    navigator.clipboard.writeText($codeForCopy).catch(console.error);
}