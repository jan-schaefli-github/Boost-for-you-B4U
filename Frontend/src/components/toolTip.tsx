import React from 'react';
import '../assets/css/tooltip.css';

interface TooltipProps {
  text: string;
  position?: 'top' | 'bottom' | 'left' | 'right';
  children: React.ReactElement;
}

function Tooltip({ text, position = 'top', children }: TooltipProps) {
  const getPositionStyles = () => {
    switch (position) {
      case 'top':
        return { top: '-40px', left: '0%' };
      case 'bottom':
        return { top: '80px', left: '0%' };
      case 'left':
        return { top: '-8px', right: '120%' };
      case 'right':
        return { top: '-8px', left: '120%' };
      default:
        return { top: '-40px', left: '-50%' };
    }
  };

  return (
    <div className="hover-text">
      {React.cloneElement(children, {
        className: `${children.props.className} tooltip-trigger`,
      })}
      <span className="tooltip-text" style={getPositionStyles()}>
        {text}
      </span>
    </div>
  );
}

export default Tooltip;
