import { ButtonHTMLAttributes, forwardRef } from 'react';
import { cn } from '@/shared/utils/cn';
import { Loader2 } from 'lucide-react';
import { Ripple } from './Ripple';

export interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: 'primary' | 'secondary' | 'danger' | 'ghost';
  size?: 'sm' | 'md' | 'lg';
  isLoading?: boolean;
}

export const Button = forwardRef<HTMLButtonElement, ButtonProps>(
  (
    { className, variant = 'primary', size = 'md', isLoading, children, disabled, ...props },
    ref
  ) => {
    const baseStyles =
      'relative overflow-hidden inline-flex items-center justify-center font-medium transition-all duration-300 ease-out focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none rounded-xl active:scale-[0.98]';

    const variants = {
      primary: 'bg-[#007AFF] text-white hover:bg-blue-600',
      secondary: 'bg-zinc-200 dark:bg-[#2c2c2e] text-zinc-900 dark:text-zinc-50 hover:bg-zinc-300 dark:hover:bg-[#3a3a3c]',
      danger: 'bg-[#FF3B30] text-white hover:bg-red-600',
      ghost: 'bg-transparent text-zinc-600 dark:text-zinc-400 hover:text-zinc-900 dark:hover:text-zinc-50 hover:bg-zinc-100 dark:hover:bg-[#2c2c2e]',
    };

    const sizes = {
      sm: 'h-9 px-3 text-sm',
      md: 'h-12 px-4 text-base rounded-2xl',
      lg: 'h-14 px-8 text-lg rounded-[20px]',
    };

    return (
      <button
        ref={ref}
        className={cn(baseStyles, variants[variant], sizes[size], className)}
        disabled={isLoading || disabled}
        {...props}
      >
        {isLoading && <Loader2 className="w-4 h-4 mr-2 animate-spin" />}
        {children}
        <Ripple color={variant === 'ghost' ? 'rgba(0,0,0,0.1)' : 'rgba(255, 255, 255, 0.3)'} />
      </button>
    );
  }
);

Button.displayName = 'Button';
