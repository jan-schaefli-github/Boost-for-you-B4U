import React, { useState, useEffect } from 'react';
import '../assets/css/tooltip.css';

interface TooltipProps {
  text: string;
  position?: { top: string; left: string };
  children: React.ReactElement;
}

function Tooltip({ text, position = { top: '-45px', left: '20%' }, children }: TooltipProps) {
  const [showTooltip, setShowTooltip] = useState(false);

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
      <div className="hover-text" onMouseOver={handleTooltipToggle}>
        {React.cloneElement(children, {
          className: `${children.props.className} tooltip-trigger`,
        })}
        {showTooltip && (
          <span className="tooltip-text" style={position}>
            {text}
          </span>
        )}
      </div>
    </div>
  );
}

export default Tooltip;
