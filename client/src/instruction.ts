export function connectInstructionToggle() {
    const $button = document.querySelector(".js-instruction-toggle-button");
    const $state = document.querySelector(".js-instruction-visibility-state");
    const $instruction = document.querySelector(".js-instruction");

    $button.addEventListener("click", () => {
        $instruction.classList.toggle("hide");
        $state.classList.toggle("active");
    });
}
