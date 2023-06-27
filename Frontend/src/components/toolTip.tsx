import React, { useState, useEffect } from 'react';
import '../assets/css/tooltip.css';

interface TooltipProps {
  text: string;
  position?: 'top' | 'bottom' | 'left' | 'right';
  children: React.ReactElement;
}

function Tooltip({ text, position = 'top', children }: TooltipProps) {
  const [showTooltip, setShowTooltip] = useState(false);

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

  useEffect(() => {
    let timeoutId: ReturnType<typeof setTimeout>;

    if (showTooltip) {
      timeoutId = setTimeout(() => {
        setShowTooltip(false);
      }, 2000);
    }

    return () => {
      if (timeoutId) {
        clearTimeout(timeoutId);
      }
    };
  }, [showTooltip]);

  const handleTooltipToggle = () => {
    setShowTooltip(!showTooltip);
  };

  return (
    <div className='tool-tip'>
    <div className="hover-text" onClick={handleTooltipToggle}>
      {React.cloneElement(children, {
        className: `${children.props.className} tooltip-trigger`,
      })}
      {showTooltip && (
        <span className="tooltip-text" style={getPositionStyles()}>
          {text}
        </span>
      )}
    </div>
    </div>
  );
}

export default Tooltip;
