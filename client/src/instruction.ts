export function initInstruction() {
    initToggle();
    updateInstructionRepositoryExistsStatus();
}

function initToggle() {
    const $button = document.querySelector(".js-instruction-toggle-button");
    const $instruction = document.querySelector(".js-instruction");

    $button.addEventListener("click", () => {
        $instruction.classList.toggle("active");
    });
}

function updateInstructionRepositoryExistsStatus() {
    const $step = document.querySelector(".js-instruction-repository-exists");
    if ($step.classList.contains("active")) {
        return;
    }

    const username = document.body.getAttribute("data-session-profile-username");
    if (username === "") {
        return;
    }

    fetch(`https://api.github.com/repos/${username}/${username}`)
        .then(function (response) {
            $step.classList.add("active")
        })
        .catch(console.error);
}
