let chart;
let polarity;

(function ($) {

	var $window = $(window),
		$body = $('body'),
		$main = $('#main');

	// Breakpoints.
	breakpoints({
		xlarge: ['1281px', '1680px'],
		large: ['981px', '1280px'],
		medium: ['737px', '980px'],
		small: ['481px', '736px'],
		xsmall: ['361px', '480px'],
		xxsmall: [null, '360px']
	});

	// Play initial animations on page load.
	$window.on('load', function () {
		window.setTimeout(function () {
			$body.removeClass('is-preload');
		}, 100);
	});

	// Nav.
	var $nav = $('#nav');

	if ($nav.length > 0) {

		// Shrink effect.
		$main
			.scrollex({
				mode: 'top',
				enter: function () {
					$nav.addClass('alt');
				},
				leave: function () {
					$nav.removeClass('alt');
				},
			});

		// Links.
		var $nav_a = $nav.find('a');

		try {
			$nav_a
				.scrolly({
					speed: 1000,
					offset: function () {
						return $nav.height();
					}
				})
				.on('click', function () {

					var $this = $(this);

					// External link? Bail.
					if ($this.attr('href').charAt(0) != '#')
						return;

					// Deactivate all links.
					$nav_a
						.removeClass('active')
						.removeClass('active-locked');

					// Activate link *and* lock it (so Scrollex doesn't try to activate other links as we're scrolling to this one's section).
					$this
						.addClass('active')
						.addClass('active-locked');

				})
				.each(function () {

					var $this = $(this),
						id = $this.attr('href'),
						$section = $(id);

					// No section for this link? Bail.
					if ($section.length < 1)
						return;

					// Scrollex.
					$section.scrollex({
						mode: 'middle',
						initialize: function () {

							// Deactivate section.
							if (browser.canUse('transition'))
								$section.addClass('inactive');

						},
						enter: function () {

							// Activate section.
							$section.removeClass('inactive');

							// No locked links? Deactivate all links and activate this section's one.
							if ($nav_a.filter('.active-locked').length == 0) {

								$nav_a.removeClass('active');
								$this.addClass('active');

							}

							// Otherwise, if this section's link is the one that's locked, unlock it.
							else if ($this.hasClass('active-locked'))
								$this.removeClass('active-locked');

						}
					});

				});
		} catch (err) {}
	}

	// Scrolly.
	$('.scrolly').scrolly({
		speed: 1000
	});

	let limit = 1000;

	let x = 20;
	let data = [];
	let dataSeries = {
		type: "line"
	};
	let dataPoints = [];
	let eq_settings = "";
	
	// $.getJSON("./assets/datadir/eq-settings.json", function (eQdata) {
	// $.ajax({
	// 	url: "./assets/datadir/eq-settings.json",
	// 	async: false,
	// 	success: function (eQdata) {
	// 		eq_settings = eQdata;
	// 		$.each(eQdata, function (key, value) {
	// 			dataPoints.push({
	// 				x: eq_settings[key].frequency,
	// 				y: eq_settings[key].gain
	// 			});
	// 		});
	// 	}
	// });
	// dataSeries.dataPoints = dataPoints;
	// data.push(dataSeries);

	var options = {
		animationEnabled: true,
		//zoomEnabled: true,
		//zoomType: "xy",
		exportEnabled: true,
		title: {
			text: "Equalizer Setup"
		},
		theme: "light2",
		annotation: {
			annotations: [
				{
					type: 'line',
					mode: 'horizontal',
					scaleID: 'y-axis-0',
					value: 25,
					draggable: true,
					onDrag: function(event) {
						console.log(event.subject.config.value);
					}
				}
			]
		},
		axisX: {
			//logarithmic: true,
			title: "Band",
			minimum: 0,
			//valueFormatString:" ",
			stripLines: [{
					value: 10,
				},
				{
					value: 20,
				},
				{
					value: 30,
				},
				{
					value: 40,
				}
			]
		},
		axisY: {
			includeZero: false,
			minimum: -30,
			maximum: 30,
			title: "Gain (dB)"
		},
		//data: data
		data: [{
			type: "line",
      		markerSize: 2,
      		lineThickness: 3,
			dataPoints: [{
				x: 0,
				y: 0
			},
			{
				x: 1,
				y: 0
			},
			{
				x: 2,
				y: 0
			},
			{
				x: 3,
				y: 0
			},
			{
				x: 4,
				y: 0
			},
			{
				x: 5,
				y: 0
			},
			{
				x: 6,
				y: 0
			},
			{
				x: 7,
				y: 0
			},
			{
				x: 8,
				y: 0
			},
			{
				x: 9,
				y: 0
			},
			{
				x: 10,
				y: 0
			},
			{
				x: 11,
				y: 0
			},
			{
				x: 12,
				y: 0
			},
			{
				x: 13,
				y: 0
			},
			{
				x: 14,
				y: 0
			},
			{
				x: 15,
				y: 0
			},
			{
				x: 16,
				y: 0
			},
			{
				x: 17,
				y: 0
			},
			{
				x: 18,
				y: 0
			},
			{
				x: 19,
				y: 0
			},
			{
				x: 20,
				y: 0
			},
			{
				x: 21,
				y: 0
			},
			{
				x: 22,
				y: 0
			},
			{
				x: 23,
				y: 0
			},
			{
				x: 24,
				y: 0
			},
			{
				x: 25,
				y: 0
			},
			{
				x: 26,
				y: 0
			},
			{
				x: 27,
				y: 0
			},
			{
				x: 28,
				y: 0
			},
			{
				x: 29,
				y: 0
			},
			{
				x: 30,
				y: 0
			},
			{
				x: 31,
				y: 0
			},
			{
				x: 32,
				y: 0
			},
			{
				x: 33,
				y: 0
			},
			{
				x: 34,
				y: 0
			},
			{
				x: 35,
				y: 0
			},
			{
				x: 36,
				y: 0
			},
			{
				x: 37,
				y: 0
			},
			{
				x: 38,
				y: 0
			},
			{
				x: 39,
				y: 0
			},
			{
				x: 40,
				y: 0
			}]
		}]
	};

	//$("#chartContainer").CanvasJSChart(options);
	chart = new CanvasJS.Chart("chartContainer", options);
		chart.render();

		var xSnapDistance = 2 / 2;
		var ySnapDistance = 2;

		var xValue, yValue;

		var mouseDown = false;
		var selected = null;
		var changeCursor = false;

		var timerId = null;

		$("#chartContainer > .canvasjs-chart-container").on("mousedown", function(e) {
			mouseDown = true;
			getPosition(e);
			searchDataPoint();
		});

		$("#chartContainer > .canvasjs-chart-container").on("mouseup", function(e) {
			if (selected != null) {
				chart.data[0].dataPoints[selected].y = yValue;
				chart.render();
				mouseDown = false;
			}
		});

		$("#chartContainer > .canvasjs-chart-container").on("mousemove", function(e) {
			getPosition(e);
			if (mouseDown) {
				clearTimeout(timerId);
				timerId = setTimeout(function() {
					if (selected != null) {
					chart.data[0].dataPoints[selected].y = yValue;
					chart.render();
					}
				}, 0);
			} else {
				searchDataPoint();
				if (changeCursor) {
					chart.data[0].set("cursor", "n-resize");
				} else {
					chart.data[0].set("cursor", "default");
				}
			}
		});


		function getPosition(e) {
			var parentOffset = $("#chartContainer > .canvasjs-chart-container").offset();
			var relX = e.pageX - parentOffset.left;
			var relY = e.pageY - parentOffset.top;
			xValue = Math.round(chart.axisX[0].convertPixelToValue(relX));
			yValue = Math.round(chart.axisY[0].convertPixelToValue(relY));
		}

		function searchDataPoint() {
			var dps = chart.data[0].dataPoints;
			for (var i = 0; i < dps.length; i++) {
			  	if ((xValue >= dps[i].x - xSnapDistance && xValue <= dps[i].x + xSnapDistance) && (yValue >= dps[i].y - ySnapDistance && yValue <= dps[i].y + ySnapDistance)) {
					if (mouseDown) {
						selected = i;
						break;
					} else {
						changeCursor = true;
						break;
					}
			  	} else {
					selected = null;
					changeCursor = false;
			  	}
			}
		}

	try {
		
	} catch (err) {

	}

	$('.btn-toggle').click(function () {
		$(this).find('.btn').toggleClass('active');
		$(this).find('.btn').toggleClass('btn-default');
	});

})(jQuery);


// function openCity(evt, cityName) {
// 	var i, tabcontent, tablinks;
// 	tabcontent = document.getElementsByClassName("tabcontent");
// 	for (i = 0; i < tabcontent.length; i++) {
// 		tabcontent[i].style.display = "none";
// 	}
// 	tablinks = document.getElementsByClassName("tablinks");
// 	for (i = 0; i < tablinks.length; i++) {
// 		tablinks[i].className = tablinks[i].className.replace(" active", "");
// 	}
// 	document.getElementById(cityName).style.display = "block";
// 	evt.currentTarget.className += " active";
// }

// Get the element with id="defaultOpen" and click on it
//document.getElementById("defaultOpen").click();


class Switch {
	constructor(
		selector, {
			highlightClass = "c-switch__highlight",
			activeClass = "is-active"
		} = {}) {
		this.activeClass = activeClass;
		this.element = document.querySelector(selector);
		this.buttons = this.element.querySelectorAll("button");
		this.highlight = this.element.querySelector(`.${highlightClass}`);
		this.activeBtn = this.element.querySelector(`button.${this.activeClass}`);

		if (!this.activeBtn) {
			this.activeBtn = this.buttons[0];
			this.activeBtn.classList.add(this.activeClass);
		}

		this._highlight();
		this._addEvents();
	}

	_highlight() {
		const w = this.activeBtn.offsetWidth;
		const x = this.activeBtn.offsetLeft;

		this.highlight.style.width = `${w}px`;
		this.highlight.style.transform = `translateX(${x}px)`;
	}

	_dispatchEvent() {
		this.element.dispatchEvent(
			new CustomEvent("change", {
				detail: this.activeBtn.dataset.value
			}));

	}

	_addEvents() {
		[].forEach.call(this.buttons, (btn) =>
			btn.addEventListener("click", e => {
				if (this.activeBtn === e.target) return;

				this.activeBtn.classList.remove(this.activeClass);
				this.activeBtn = e.target;
				this.activeBtn.classList.add(this.activeClass);

				this._highlight();
				this._dispatchEvent();
			}));

	}
}
const mySwitch = new Switch(".c-switch");
mySwitch.element.addEventListener("change", e => {
	console.log(e.detail);
	if (e.detail === "first") {
		console.log("0");
		polarity = "0";
	}
	else {
		console.log("1");
		polarity = "1";
	}
});


$(function() {
  	$('input[type=file]').change(function(){
		var t = $(this).val();
		var labelText = 'File : ' + t.substr(12, t.length);
		$(this).prev('label').text(labelText);
  	})
});

function saveEQSettings() {
	let settings = "";

	// Save as csv
	for(let i=0; i <=40; i++) {
		//console.log(chart.axisX[0].dataSeries[0].dataPoints[i].y);
		settings += (chart.axisX[0].dataSeries[0].dataPoints[i].y.toString() + ", ")
	}

	settings += $("input[type=range]").val().toString() + ", ";

	settings += polarity + ", ";

	settings += $("#speakerList").find(":selected").text().toString();

	let jsonObj = {
		"settings": settings
	};

	console.log(JSON.stringify(jsonObj));
	
	$.ajax({
		type: 'POST',
		url: '/saveConfig',
		data: JSON.stringify(jsonObj),
		contentType: "application/json",
		dataType: 'json',
		success: function(retVal) {
			alert('Success: ' + JSON.stringify(retVal));
		},
		error: function(err) {
			console.log('Error: ' + JSON.stringify(err));
		}
	});
}