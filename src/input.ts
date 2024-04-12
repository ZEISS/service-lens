import 'htmx.org'
import 'hyperscript.org'

customElements.define(
  'chat-input',
  class extends HTMLFormElement {
    connectedCallback() {
      if (!this.isConnected) return

      this.dataset.hxGet = '/'
      this.dataset.hxTarget = '#messages'
      this.dataset.hxSwap = 'afterbegin'

      this.addEventListener('htmx:beforeSend', () => this.reset())
    }
  },
  { extends: 'form' }
)
