//
// Taken from github.com/carbon-design-system/carbon/packages/react/src/internal/keyboard/navigation.js
//

// add all the elements inside modal which you want to make focusable
const selectorTabbable: string = `
  a[href], area[href], input:not([disabled]):not([tabindex='-1']),
  button:not([disabled]):not([tabindex='-1']),select:not([disabled]):not([tabindex='-1']),
  textarea:not([disabled]):not([tabindex='-1']),
  iframe, object, embed, *[tabindex]:not([tabindex='-1']):not([disabled]), *[contenteditable=true]
`;

// Assuming you have some type for the Action, replace 'any' with the appropriate type
export default function focusTrap(node: HTMLElement): { destroy: () => void } {
	const handleFocusTrap = (e: KeyboardEvent): void => {
		const isTabPressed: boolean = e.key === 'Tab' || e.keyCode === 9;

		if (!isTabPressed) {
			return;
		}

		const tabbable: HTMLElement[] = Array.from(node.querySelectorAll(selectorTabbable));

		let index: number = tabbable.indexOf((document.activeElement as HTMLElement) ?? node);
		if (index === -1 && e.shiftKey) index = 0;
		index += tabbable.length + (e.shiftKey ? -1 : 1);
		index %= tabbable.length;
		tabbable[index].focus();

		e.preventDefault();
	};

	document.addEventListener('keydown', handleFocusTrap, true);

	return {
		destroy() {
			document.removeEventListener('keydown', handleFocusTrap, true);
		}
	};
}
