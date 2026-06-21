import { InputHTMLAttributes, forwardRef } from 'react';
import { cn } from '@/shared/utils/cn';

export interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  label?: string;
  error?: string;
}

export const Input = forwardRef<HTMLInputElement, InputProps>(
  ({ className, label, error, ...props }, ref) => {
    return (
      <div className="flex flex-col gap-1.5 w-full">
        {label && (
          <label className="text-sm font-medium text-zinc-500 dark:text-zinc-400 ml-1">
            {label}
          </label>
        )}
        <input
          ref={ref}
          className={cn(
            'flex h-12 w-full rounded-2xl bg-white dark:bg-[#1c1c1e] border border-zinc-200 dark:border-white/5 px-4 py-2 text-base text-zinc-900 dark:text-zinc-50 shadow-sm transition-all duration-300',
            'file:border-0 file:bg-transparent file:text-sm file:font-medium',
            'placeholder:text-zinc-400 dark:placeholder:text-zinc-600',
            'focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-blue-500 focus-visible:border-transparent',
            'disabled:cursor-not-allowed disabled:opacity-50',
            error && 'border-[#FF3B30] focus-visible:ring-[#FF3B30]',
            className
          )}
          {...props}
        />
        {error && <span className="text-sm text-[#FF3B30] ml-1">{error}</span>}
      </div>
    );
  }
);

Input.displayName = 'Input';
