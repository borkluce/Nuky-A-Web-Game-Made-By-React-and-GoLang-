import React from "react"

interface CInputFieldProps {
    label: string
    placeholder: string
    value: string
    onChange: (e: React.ChangeEvent<HTMLInputElement>) => void
}

const CInputField: React.FC<CInputFieldProps> = ({
    label,
    placeholder,
    value,
    onChange,
}) => {
    return (
        <div className="">
            <label
                className="block text-sm font-bold mb-2 text-black"
                htmlFor={label}
            >
                {label}
            </label>
            <input
                placeholder={placeholder}
                id={`input-${label}`}
                type="text"
                value={value}
                onChange={onChange}
                className="appearance-none border border-black border-opacity-40 w-full 
                py-3 px-5 leading-tight focus:outline-none focus:shadow-outline
                text-sm placeholder:text-black placeholder:text-opacity-40
                focus:border-2 focus:border-primary
                "
            />
        </div>
    )
}

export default CInputField
