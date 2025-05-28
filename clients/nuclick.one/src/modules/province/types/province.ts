export type Province = {
    ID: string
    province_name: string
    province_color_hex: string
    attack_count: number
    support_count: number
    destroyment_round: number
}

// --------------------------------------------------------------------

// Utility to generate a random hex color
const getRandomColor = () =>
    "#" +
    Math.floor(Math.random() * 16777215)
        .toString(16)
        .padStart(6, "0")

// Utility to generate random provinces
export const generateRandomProvinces = (count: number): Province[] => {
    const provinces: Province[] = []
    for (let i = 1; i <= count; i++) {
        provinces.push({
            ID: `${i}`,
            province_name: `Province ${i}`,
            province_color_hex: getRandomColor(),
            attack_count: Math.floor(Math.random() * 10),
            support_count: Math.floor(Math.random() * 10),
            destroyment_round: -1,
        })
    }
    return provinces
}
