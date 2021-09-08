#!/bin/bash

function CallSpider() {
	echo "- Get securityGroup in ${MCIRRegionName}"

	curl -H "${AUTH}" -sX GET http://$SpiderServer/spider/securitygroup/${CONN_CONFIG[$INDEX,$REGION]}-${POSTFIX} -H 'Content-Type: application/json' -d '{ "ConnectionName": "'${CONN_CONFIG[$INDEX,$REGION]}'"}' | jq ''
}

#function spider_get_securityGroup() {

	echo "####################################################################"
	echo "## 4. securityGroup: Get"
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

				CallSpider

			done

		done
		wait

	else
		echo ""
		
		MCIRRegionName=${CONN_CONFIG[$INDEX,$REGION]}

		CallSpider

	fi

#}

#spider_get_securityGroup