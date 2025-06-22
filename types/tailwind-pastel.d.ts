// TypeScript definitions for custom Tailwind pastel colors
// These match the existing QueueChat component's color scheme

declare module 'tailwindcss/lib/util/flattenColorPalette' {
  const flattenColorPalette: (colors: any) => any;
  export = flattenColorPalette;
}

// Pastel color definitions for Tailwind CSS
export interface PastelColors {
  'pastel-pink': string;
  'pastel-rose': string;
  'pastel-purple': string;
  'pastel-blue': string;
  'pastel-sky': string;
  'pastel-aqua': string;
  'pastel-teal': string;
  'pastel-green': string;
  'pastel-mint': string;
  'pastel-sage': string;
  'pastel-yellow': string;
  'pastel-lemon': string;
}

// Default pastel color values
export const PASTEL_COLORS: PastelColors = {
  'pastel-pink': '#FFE4E6',
  'pastel-rose': '#FFE4E1',
  'pastel-purple': '#F3E8FF',
  'pastel-blue': '#E0F2FE',
  'pastel-sky': '#E0F7FA',
  'pastel-aqua': '#E0F4F3',
  'pastel-teal': '#E6FFFA',
  'pastel-green': '#F0FDF4',
  'pastel-mint': '#ECFDF5',
  'pastel-sage': '#F0F9F4',
  'pastel-yellow': '#FEFCE8',
  'pastel-lemon': '#FFFBEB',
};

// Extend Tailwind types
declare module 'tailwindcss/tailwind-config' {
  interface TailwindConfig {
    theme?: {
      extend?: {
        colors?: PastelColors;
      };
    };
  }
}

// CSS class name helpers
export type PastelBackgroundClass = 
  | 'bg-pastel-pink'
  | 'bg-pastel-rose'
  | 'bg-pastel-purple'
  | 'bg-pastel-blue'
  | 'bg-pastel-sky'
  | 'bg-pastel-aqua'
  | 'bg-pastel-teal'
  | 'bg-pastel-green'
  | 'bg-pastel-mint'
  | 'bg-pastel-sage'
  | 'bg-pastel-yellow'
  | 'bg-pastel-lemon';

export type PastelBorderClass =
  | 'border-pastel-pink'
  | 'border-pastel-rose'
  | 'border-pastel-purple'
  | 'border-pastel-blue'
  | 'border-pastel-sky'
  | 'border-pastel-aqua'
  | 'border-pastel-teal'
  | 'border-pastel-green'
  | 'border-pastel-mint'
  | 'border-pastel-sage'
  | 'border-pastel-yellow'
  | 'border-pastel-lemon';

export type PastelTextClass =
  | 'text-pastel-pink'
  | 'text-pastel-rose'
  | 'text-pastel-purple'
  | 'text-pastel-blue'
  | 'text-pastel-sky'
  | 'text-pastel-aqua'
  | 'text-pastel-teal'
  | 'text-pastel-green'
  | 'text-pastel-mint'
  | 'text-pastel-sage'
  | 'text-pastel-yellow'
  | 'text-pastel-lemon';

export type PastelGradientClass =
  | 'from-pastel-pink'
  | 'from-pastel-rose'
  | 'from-pastel-purple'  
  | 'from-pastel-blue'
  | 'from-pastel-sky'
  | 'from-pastel-aqua'
  | 'from-pastel-teal'
  | 'from-pastel-green'
  | 'from-pastel-mint'
  | 'from-pastel-sage'
  | 'from-pastel-yellow'
  | 'from-pastel-lemon'
  | 'to-pastel-pink'
  | 'to-pastel-rose'
  | 'to-pastel-purple'
  | 'to-pastel-blue'
  | 'to-pastel-sky'
  | 'to-pastel-aqua'
  | 'to-pastel-teal'
  | 'to-pastel-green'
  | 'to-pastel-mint'
  | 'to-pastel-sage'
  | 'to-pastel-yellow'
  | 'to-pastel-lemon';