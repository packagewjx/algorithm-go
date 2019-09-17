package study

import (
	"fmt"
	"testing"
)

func TestGetEditDistance(t *testing.T) {
	println(getEditDistance("mt", "mt"))
}

func TestMy(t *testing.T) {
	words := []string{
		"all-asset",
		"asset-growth",
		"asset-growth-3year",
		"brand-asset",
		"brand-income-growth",
		"brand-infringement",
		"brand-output",
		"brand-output-profit",
		"brand-overseas",
		"brand-profit",
		"brand-profit-growth",
		"earning-share",
		"employee-level-junior",
		"employee-level-senior",
		"employee-level-university",
		"employee-number",
		"employment-contribution",
		"fan-number",
		"fixed-asset",
		"fixed-asset-new",
		"fixed-asset-original",
		"net-asset",
		"net-asset-growth",
		"net-asset-rate",
		"net-profit-growth",
		"output",
		"profit",
		"rd-institution-expense",
		"registered-asset",
		"salary",
		"salary-growth",
		"sale-growth",
		"sale-growth-3year",
		"stock-growth",
		"stock-growth-last",
		"stock-growth-this",
		"tax",
		"tax-paid-rate",
		"tax-rate",
		"turnover",
		"turnover-growth",
	}

	dict := []string{
		"abstract",
		"advertising",
		"advertising-area",
		"advertising-expense",
		"advertising-level",
		"annual-brand-support-number",
		"annual-product-placement",
		"annual-team-support",
		"annual-title-sponsor",
		"annual-tv-show-support",
		"asset-growth",
		"baidu-index",
		"brand-award",
		"brand-awareness",
		"brand-fan-number",
		"brand-favorite",
		"brand-marketing-and-public-welfare",
		"brand-media",
		"brand-media-article-number",
		"brand-news-read-count",
		"brand-premium",
		"brand-rank",
		"business-area",
		"business-kind-number",
		"cannes-lions",
		"celebrity",
		"celebrity-expense",
		"celebrity-level",
		"celebrity-number",
		"channel-cover-area",
		"channel-coverage",
		"channel-growth",
		"channel-number",
		"channel-offline-sales",
		"channel-offline-sales-growth-rate",
		"channel-offline-sales-ratio",
		"channel-online-sales",
		"channel-online-sales-growth-rate",
		"channel-online-sales-ratio",
		"channel-power",
		"channel-quality",
		"channel-sales",
		"channel-sales-growth-rate",
		"channel-sales-ratio",
		"channel-score",
		"channel-type",
		"china-4a",
		"china-500-rank",
		"china-advertising-great-wall-awards",
		"china-content-marketing-awards",
		"china-effie-award",
		"computer-rate",
		"develop-growth",
		"develop-power",
		"develop-score",
		"direct-sale-store-number",
		"doctor-number",
		"donation",
		"donation-rate",
		"earning-per-share",
		"employee-above-bachelor-ratio",
		"employee-average-salary",
		"employee-contract-rate",
		"employee-hurt-rate",
		"employee-junior-ratio",
		"employee-level-of-education",
		"employee-relation",
		"employee-salary-grow-rate",
		"employee-senior-ratio",
		"employment-contribution",
		"enterprise-growth",
		"enterprise-income-tax",
		"environment-relation",
		"environmental-protection-fund-growth-rate",
		"environmental-protection-fund-rate",
		"environmental-protection-funds",
		"export-amount",
		"fixed-asset",
		"fixed-asset-new-rate",
		"fixed-asset-rate",
		"franchise-store-number",
		"global-brand-500-rank",
		"global-fortune-500-rank",
		"golden-mouse-award",
		"government-fund",
		"government-relation",
		"has-company-culture",
		"has-company-publication",
		"has-employee-training",
		"has-manage-department",
		"has-manage-expert",
		"has-mergers-and-acquisitions",
		"has-publish-csr",
		"has-publish-quality-credit-report",
		"has-social-accountability",
		"has-standard",
		"highest-government-award",
		"industry-academia-research",
		"industry-academia-research-city",
		"industry-academia-research-country",
		"industry-academia-research-other",
		"industry-academia-research-province",
		"industry-academia-research-world",
		"industry-exhibition",
		"innovation-ability",
		"innovation-expense",
		"intellectual-property",
		"interbrand-rank",
		"internet-search-index",
		"inventory-value",
		"is-high-tech",
		"legal-dispute-number",
		"london-international-award",
		"longxi-award",
		"main-business-revenue-growth-rate",
		"manage-power",
		"manage-score",
		"market-coverage",
		"market-international",
		"market-power",
		"market-premium",
		"market-score",
		"market-share",
		"market-stable",
		"marketing-award",
		"marketing-expense",
		"media-article",
		"media-channel",
		"mobile-internet-system-rate",
		"negative-article-rate",
		"net-asset",
		"net-asset-growth-rate",
		"net-profit-growth-rate",
		"new-york-festival-award",
		"official-media",
		"official-media-article-number",
		"operation-manage",
		"organization-manage",
		"other-media",
		"other-media-article-number",
		"overseas-sale-rate",
		"patent-applied",
		"patent-invest-over-100000",
		"patent-kind",
		"patent-kind-design",
		"patent-kind-invent",
		"patent-kind-utility-model",
		"pct-application",
		"produce-ability",
		"product-level",
		"product-power",
		"product-praise-number",
		"product-qualified-change-rate",
		"product-reputation",
		"product-score",
		"profit-growth-rate",
		"profit-rate",
		"protect-ability",
		"public-relations",
		"public-relations-expense-rate",
		"public-welfare-relation",
		"rd-expense",
		"rd-institution",
		"rd-institution-improve",
		"rd-institution-independence",
		"rd-institution-number",
		"rd-institution-number-city",
		"rd-institution-number-country",
		"rd-institution-number-other",
		"rd-institution-number-province",
		"rd-institution-number-world",
		"rd-institution-outside",
		"rd-investment",
		"rd-organization-management",
		"rd-personnel-number",
		"registered-asset",
		"relation-power",
		"relation-score",
		"return-to-asset-rate",
		"revenue-growth-rate",
		"roe",
		"roi-festival-award",
		"self-media",
		"self-media-article-number",
		"self-media-fan-count",
		"self-media-number",
		"self-media-operation",
		"self-media-read-count",
		"sell-growth",
		"sell-growth-rate",
		"sell-growth-rate-3-year",
		"share-growth-rate",
		"shareholders-relation",
		"social-fan-number",
		"social-relation",
		"spread-activity",
		"spread-coverage",
		"spread-power",
		"spread-precision",
		"spread-score",
		"store-praise-rate",
		"tax-paid-rate",
		"tax-rate",
		"tiger-roar-award",
		"total-asset",
		"total-asset-growth-rate",
		"total-asset-growth-rate-3-year",
		"total-export",
		"total-export-profit-rate",
		"total-profit",
		"total-revenue",
		"trademark-honor",
		"training-expense",
		"transform-in-3-year",
		"transformation-ability",
		"value-power",
		"value-score",
	}

	for i := 0; i < len(words); i++ {
		distances := GetEditDistances(words[i], dict)[:10]
		fmt.Printf("与%s距离最近的是：\n", words[i])
		for j := 0; j < len(distances); j++ {
			fmt.Printf("%s: %d\n", distances[j].word, distances[j].distance)
		}
	}
}