import React, { useEffect } from "react"
import { useProvince } from "../../province/hooks/useProvince"

type PanelProps = {
  className?: string
}

const TopAllProvincesPanel: React.FC<PanelProps> = ({ className = "" }) => {
  // Get province data and methods from the store
  const { provinceList, isLoading, error, getAllProvinces } = useProvince()

  // Fetch all provinces on component mount
  useEffect(() => {
    getAllProvinces()
  }, [getAllProvinces])

  return (
    <div className={`p-4 ${className}`}>
      <h2 className="text-black text-lg font-bold mb-4">
        All Provinces
      </h2>

      {isLoading && <p className="text-gray-500">Loading provinces...</p>}
      
      {error && <p className="text-red-500">{error}</p>}
      
      {!isLoading && provinceList && provinceList.length === 0 && (
        <p className="text-gray-500">No province data available</p>
      )}

      {!isLoading && provinceList && provinceList.length > 0 && (
        <div className="space-y-2 max-h-[calc(100vh-180px)] overflow-y-auto pr-2">
          {provinceList.map((province) => (
            <div 
              key={province.ID} 
              className="bg-white/70 p-2 rounded-md shadow-sm border border-gray-200"
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

export default TopAllProvincesPanel