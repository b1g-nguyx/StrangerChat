"use client";

import { useState, useEffect } from "react";
import { motion, AnimatePresence } from "framer-motion";
import { Globe, Shield, Sparkles } from "lucide-react";

const slides = [
  {
    id: 1,
    title: "Meet The World",
    description: "Instantly connect with random strangers globally. Discover new cultures and make friends.",
    icon: <Globe className="w-12 h-12 text-[#007AFF]" />,
    color: "bg-blue-500/10 text-blue-500",
  },
  {
    id: 2,
    title: "100% Anonymous",
    description: "Your privacy is our priority. No real names required. Chat freely and safely.",
    icon: <Shield className="w-12 h-12 text-emerald-500" />,
    color: "bg-emerald-500/10 text-emerald-500",
  },
  {
    id: 3,
    title: "Beautifully Simple",
    description: "An elegant, distraction-free interface designed to keep the focus on your conversations.",
    icon: <Sparkles className="w-12 h-12 text-purple-500" />,
    color: "bg-purple-500/10 text-purple-500",
  },
];

export const HeroSlider = () => {
  const [currentSlide, setCurrentSlide] = useState(0);

  useEffect(() => {
    const timer = setInterval(() => {
      setCurrentSlide((prev) => (prev + 1) % slides.length);
    }, 4000);
    return () => clearInterval(timer);
  }, []);

  return (
    <div className="relative w-full max-w-2xl mx-auto h-[320px] mb-8 overflow-hidden rounded-[32px] bg-white/40 dark:bg-[#1c1c1e]/40 backdrop-blur-xl border border-black/5 dark:border-white/10 shadow-[0_8px_30px_rgb(0,0,0,0.04)] dark:shadow-[0_8px_30px_rgb(0,0,0,0.12)]">
      <AnimatePresence mode="wait">
        <motion.div
          key={currentSlide}
          initial={{ opacity: 0, x: 50, scale: 0.95 }}
          animate={{ opacity: 1, x: 0, scale: 1 }}
          exit={{ opacity: 0, x: -50, scale: 0.95 }}
          transition={{ duration: 0.5, ease: "easeOut" }}
          className="absolute inset-0 flex flex-col items-center justify-center p-8 text-center"
        >
          <div className={`w-24 h-24 rounded-full flex items-center justify-center mb-6 ${slides[currentSlide].color}`}>
            {slides[currentSlide].icon}
          </div>
          <h2 className="text-3xl font-bold text-zinc-900 dark:text-zinc-50 mb-4 tracking-tight">
            {slides[currentSlide].title}
          </h2>
          <p className="text-zinc-500 dark:text-zinc-400 max-w-md text-lg">
            {slides[currentSlide].description}
          </p>
        </motion.div>
      </AnimatePresence>

      {/* Pagination Dots */}
      <div className="absolute bottom-6 left-0 right-0 flex justify-center gap-2 z-10">
        {slides.map((_, index) => (
          <button
            key={index}
            onClick={() => setCurrentSlide(index)}
            className={`w-2 h-2 rounded-full transition-all duration-300 ${
              currentSlide === index
                ? "bg-[#007AFF] w-6"
                : "bg-zinc-300 dark:bg-zinc-600 hover:bg-zinc-400 dark:hover:bg-zinc-500"
            }`}
            aria-label={`Go to slide ${index + 1}`}
          />
        ))}
      </div>
    </div>
  );
};
