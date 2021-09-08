#!/bin/bash

function CallTB() {
	echo "- Lookup specs in ${MCIRRegionName}"

	resp=$(
	curl -H "${AUTH}" -sX GET http://$TumblebugServer/tumblebug/lookupSpecs -H 'Content-Type: application/json' -d @- <<EOF
		{ 
			"connectionName": "${CONN_CONFIG[$INDEX,$REGION]}"
		}
EOF
	)
	echo ${resp} | jq ''
	echo ""
}

#function lookup_spec_list() {

SECONDS=0

	echo "####################################################################"
	echo "## 7. spec: Lookup Spec List"
	echo "####################################################################"

	source ../init.sh

	if [ "${INDEX}" == "0" ]; then
		echo "[Parallel excution for all CSP regions]"

		INDEXX=${NumCSP}
		for ((cspi = 1; cspi <= INDEXX; cspi++)); do
			echo $i
			INDEXY=${NumRegion[$cspi]}
			CSP=${CSPType[$cspi]}
			for ((cspj = 1; cspj <= INDEXY; cspj++)); do
				# INDEX=$(($INDEX+1))

				echo $j
				INDEX=$cspi
				REGION=$cspj
				echo $CSP
				echo $REGION
				echo ${RegionName[$cspi,$cspj]}
				MCIRRegionName=${RegionName[$cspi,$cspj]}

				CallTB

			done

		done
		wait

	else
		echo ""
		
		MCIRRegionName=${CONN_CONFIG[$INDEX,$REGION]}

		CallTB

	fi

source ../common-functions.sh
printElapsed $@

#lookup_spec_list
