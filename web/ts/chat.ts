/*
    Returns 'chatOpen' value from localStorage or defaults with false.
    Calls 'scrollToBottom' after 250ms, so that the 'chatBox' is already
    visible.
*/
function initChat() {
    const val = window.localStorage.getItem("chatOpen");
    if (val) {
        setTimeout(scrollToBottom, 250);
    }
    return val ? JSON.parse(val) : false;
}

/*
    Scroll to the bottom of the 'chatBox'
 */
function scrollToBottom() {
    document.getElementById("chatBox").scrollTop = document.getElementById("chatBox").scrollHeight;
}

/*
    Saves negated show value in localStorage with key 'chatOpen'
    and returns the value.
    Calls 'scrollToBottom' after 250ms, so that the 'chatBox' is already
    visible.
 */
function toggleChat(show: boolean) {
    const neg = !show;
    if (neg) {
        setTimeout(scrollToBottom, 250);
    }
    window.localStorage.setItem("chatOpen", JSON.stringify(neg));
    return neg;
}