'use client';

import { useState, useLayoutEffect } from 'react';
import { motion, AnimatePresence } from 'framer-motion';

export const Ripple = ({ color = 'rgba(255, 255, 255, 0.3)', duration = 600 }) => {
  const [ripples, setRipples] = useState<{ x: number; y: number; size: number; id: number }[]>([]);

  const useIsomorphicLayoutEffect = typeof window !== 'undefined' ? useLayoutEffect : () => {};

  const addRipple = (e: React.MouseEvent<HTMLDivElement, MouseEvent>) => {
    const container = e.currentTarget.getBoundingClientRect();
    const size = Math.max(container.width, container.height);
    const x = e.clientX - container.left - size / 2;
    const y = e.clientY - container.top - size / 2;
    const newRipple = { x, y, size, id: Date.now() };

    setRipples((prevRipples) => [...prevRipples, newRipple]);
  };

  useIsomorphicLayoutEffect(() => {
    if (ripples.length > 0) {
      const timer = setTimeout(() => {
        setRipples((prevRipples) => prevRipples.slice(1));
      }, duration);
      return () => clearTimeout(timer);
    }
  }, [ripples, duration]);

  return (
    <div
      className="absolute inset-0 overflow-hidden rounded-inherit"
      onMouseDown={addRipple}
    >
      <AnimatePresence>
        {ripples.map((ripple) => (
          <motion.span
            key={ripple.id}
            initial={{ opacity: 0.5, scale: 0 }}
            animate={{ opacity: 0, scale: 2.5 }}
            exit={{ opacity: 0 }}
            transition={{ duration: duration / 1000, ease: 'easeOut' }}
            style={{
              position: 'absolute',
              top: ripple.y,
              left: ripple.x,
              width: ripple.size,
              height: ripple.size,
              backgroundColor: color,
              borderRadius: '100%',
              pointerEvents: 'none',
            }}
          />
        ))}
      </AnimatePresence>
    </div>
  );
};
