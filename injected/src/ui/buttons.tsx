import classNames from 'classnames';
import { FunctionComponent, JSX } from 'preact';
import register from 'preact-custom-element';

const buttonStyles = `
	.cs-button {
		width: auto;
		padding: 6px 12px;

		background-color: var(--cs-col-primary);
		color: white;
		text-transform: uppercase;

		border: none;
		border-radius: 2px;

		cursor: pointer;

		transition: all 300ms;
	}

	.cs-button:hover {
		background-color: var(--cs-col-primary-hover);
	}

	.cs-button:focus {
		outline: none;
	}

	.cs-button:focus-visible {
		outline: unset;
		background-color: white;
		color: black;
	}
`;

export type ButtonProps = JSX.HTMLAttributes<HTMLButtonElement>;

const Button: FunctionComponent<ButtonProps> = (props) => (
  <button className={classNames('cs-button', props.className)} {...props}>
    {props.children}
  </button>
);

export default (css: CSSStyleSheet) => {
  register(Button, 'cs-button', [], {
    shadow: true,
    adoptedStyleSheets: [css],
  });
  return buttonStyles;
};