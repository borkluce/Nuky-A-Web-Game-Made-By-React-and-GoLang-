import React, { useEffect } from "react"
import { useProvince } from "../../province/hooks/useProvince"

type PanelProps = {
  className?: string
}

const Top5ProvincesPanel: React.FC<PanelProps> = ({ className = "" }) => {
  // Get province data and methods from the store
  const { topProvinces, isLoading, error, getTopProvinces } = useProvince()

  // Fetch top provinces on component mount
  useEffect(() => {
    getTopProvinces()
  }, [getTopProvinces])

  return (
    <div className={`p-4 ${className}`}>
      <h2 className="text-black text-lg font-bold mb-4">
        Top 5 Dangerous States
      </h2>

      {isLoading && <p className="text-gray-500">Loading top provinces...</p>}
      
      {error && <p className="text-red-500">{error}</p>}
      
      {!isLoading && topProvinces && topProvinces.length === 0 && (
        <p className="text-gray-500">No province data available</p>
      )}

      {!isLoading && topProvinces && topProvinces.length > 0 && (
        <div className="space-y-3">
          {topProvinces.slice(0, 5).map((province) => (
            <div 
              key={province.ID} 
              className="bg-white/70 p-3 rounded-md shadow-sm border border-gray-200"
            >
              <p className="font-semibold">{province.province_name}</p>
              <div className="flex justify-between text-sm text-gray-600 mt-1">
                <span>Damage: {province.attack_count-province.support_count}</span>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  )
}

export default Top5ProvincesPanel
